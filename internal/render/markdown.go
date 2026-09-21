// Package render turns an inspected *model.Module into documentation output.
//
// The current renderer is a verbose, debug-oriented view that dumps every
// field of every block type. It is meant for development and troubleshooting;
// the curated, reader-facing layout is a separate renderer (planned). See
// ROADMAP.md.
package render

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/remoterabbit/open-inspector/pkg/model"
)

// Markdown renders a module as a verbose Markdown reference document.
//
// This is a debug-oriented view: every field of every block type is rendered,
// including the full source position (filename + start/end line:column:byte),
// so the output doubles as a visualization of what the inspector parsed.
func Markdown(m *model.Module) string {
	var b strings.Builder

	title := filepath.Base(m.Path)
	fmt.Fprintf(&b, "# Module `%s`\n\n", title)
	fmt.Fprintf(&b, "- **Path:** %s\n", code(m.Path))
	if len(m.RequiredCore) > 0 {
		fmt.Fprintf(&b, "- **Required core:** %s\n", code(strings.Join(m.RequiredCore, ", ")))
	}
	b.WriteString("\n")

	renderRequiredProviders(&b, m.RequiredProviders)
	renderProviders(&b, m.Providers)
	renderVariables(&b, m.Variables)
	renderOutputs(&b, m.Outputs)
	renderLocals(&b, m.Locals)
	renderResources(&b, "Managed Resources", m.ManagedResources)
	renderResources(&b, "Data Resources", m.DataResources)
	renderEphemeralResources(&b, m.EphemeralResources)
	renderModuleCalls(&b, m.ModuleCalls)
	renderMoved(&b, m.Moved)
	renderImports(&b, m.Imports)
	renderRemoved(&b, m.Removed)
	renderChecks(&b, m.Checks)
	renderSchemaFindings(&b, m)
	renderDiagnostics(&b, m.Diagnostics)

	return b.String()
}

func renderRequiredProviders(b *strings.Builder, reqs map[string]model.ProviderRequirement) {
	if len(reqs) == 0 {
		return
	}
	b.WriteString("## Required Providers\n\n")
	b.WriteString("| Name | Source | Version Constraints | Configuration Aliases | Position |\n")
	b.WriteString("|------|--------|---------------------|-----------------------|----------|\n")
	for _, name := range sortedKeys(reqs) {
		r := reqs[name]
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n",
			code(name),
			code(r.Source),
			code(strings.Join(r.VersionConstraints, ", ")),
			code(strings.Join(r.ConfigurationAliases, ", ")),
			pos(r.Position))
	}
	b.WriteString("\n")
}

func renderProviders(b *strings.Builder, providers []model.ProviderConfig) {
	if len(providers) == 0 {
		return
	}
	b.WriteString("## Providers\n\n")
	b.WriteString("| Name | Alias | For Each | Position |\n")
	b.WriteString("|------|-------|----------|----------|\n")
	for _, p := range providers {
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n",
			code(p.Name), code(p.Alias), exprCell(p.ForEach), pos(p.Position))
	}
	b.WriteString("\n")
}

func renderVariables(b *strings.Builder, vars []model.Variable) {
	if len(vars) == 0 {
		return
	}
	b.WriteString("## Inputs\n\n")
	b.WriteString("| Name | Type | Default | Required | Description | Comment | Sensitive | Nullable | Ephemeral | Validations | Position |\n")
	b.WriteString("|------|------|---------|:--------:|-------------|---------|:---------:|:--------:|:---------:|:-----------:|----------|\n")
	for _, v := range vars {
		required := "yes"
		if v.Default != nil {
			required = "no"
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %d | %s |\n",
			code(v.Name),
			typeCell(v.Type, v.TypeSpec),
			defaultCell(v.Default, v.DefaultValue),
			required,
			oneLine(v.Description),
			oneLine(v.Comment),
			boolCell(v.Sensitive),
			boolPtrCell(v.Nullable),
			boolCell(v.Ephemeral),
			len(v.Validations),
			pos(v.Position))
	}
	b.WriteString("\n")
}

func renderOutputs(b *strings.Builder, outs []model.Output) {
	if len(outs) == 0 {
		return
	}
	b.WriteString("## Outputs\n\n")
	b.WriteString("| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |\n")
	b.WriteString("|------|-------|------------|-------------|---------|:---------:|:---------:|------------|----------|\n")
	for _, o := range outs {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			code(o.Name),
			exprValueCell(o.Value),
			refsCell(o.Value.References),
			oneLine(o.Description),
			oneLine(o.Comment),
			boolCell(o.Sensitive),
			boolCell(o.Ephemeral),
			code(strings.Join(o.DependsOn, ", ")),
			pos(o.Position))
	}
	b.WriteString("\n")
}

func renderLocals(b *strings.Builder, locals []model.Local) {
	if len(locals) == 0 {
		return
	}
	b.WriteString("## Locals\n\n")
	b.WriteString("| Name | Value | References | Position |\n")
	b.WriteString("|------|-------|------------|----------|\n")
	for _, l := range locals {
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n",
			code(l.Name), exprValueCell(l.Value), refsCell(l.Value.References), pos(l.Position))
	}
	b.WriteString("\n")
}

func renderResources(b *strings.Builder, heading string, resources []model.Resource) {
	if len(resources) == 0 {
		return
	}
	fmt.Fprintf(b, "## %s\n\n", heading)
	b.WriteString("| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |\n")
	b.WriteString("|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|\n")
	for _, r := range resources {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			code(string(r.Mode)),
			code(r.Type),
			code(r.Name),
			code(r.Provider),
			exprCell(r.Count),
			exprCell(r.ForEach),
			code(strings.Join(r.DependsOn, ", ")),
			attrNamesCell(r.NestedBody),
			lifecycleCell(r.Lifecycle),
			oneLine(r.Comment),
			pos(r.Position))
	}
	b.WriteString("\n")
}

func renderEphemeralResources(b *strings.Builder, resources []model.EphemeralResource) {
	if len(resources) == 0 {
		return
	}
	b.WriteString("## Ephemeral Resources\n\n")
	b.WriteString("| Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |\n")
	b.WriteString("|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|\n")
	for _, r := range resources {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			code(r.Type),
			code(r.Name),
			code(r.Provider),
			exprCell(r.Count),
			exprCell(r.ForEach),
			code(strings.Join(r.DependsOn, ", ")),
			attrNamesCell(r.NestedBody),
			lifecycleCell(r.Lifecycle),
			oneLine(r.Comment),
			pos(r.Position))
	}
	b.WriteString("\n")
}

func renderModuleCalls(b *strings.Builder, calls []model.ModuleCall) {
	if len(calls) == 0 {
		return
	}
	b.WriteString("## Module Calls\n\n")
	b.WriteString("| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |\n")
	b.WriteString("|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|\n")
	for _, c := range calls {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			code(c.Name),
			code(c.Source),
			code(c.Version),
			exprCell(c.Count),
			exprCell(c.ForEach),
			code(strings.Join(c.DependsOn, ", ")),
			mapCell(c.Providers),
			inputsCell(c.Inputs),
			oneLine(c.Comment),
			pos(c.Position))
	}
	b.WriteString("\n")
}

func renderMoved(b *strings.Builder, blocks []model.MovedBlock) {
	if len(blocks) == 0 {
		return
	}
	b.WriteString("## Moved\n\n")
	b.WriteString("| From | To | Position |\n")
	b.WriteString("|------|----|----------|\n")
	for _, m := range blocks {
		fmt.Fprintf(b, "| %s | %s | %s |\n", code(m.From), code(m.To), pos(m.Position))
	}
	b.WriteString("\n")
}

func renderImports(b *strings.Builder, blocks []model.ImportBlock) {
	if len(blocks) == 0 {
		return
	}
	b.WriteString("## Imports\n\n")
	b.WriteString("| To | ID | Provider | Position |\n")
	b.WriteString("|----|----|----------|----------|\n")
	for _, im := range blocks {
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n",
			code(im.To), exprValueCell(im.ID), code(im.Provider), pos(im.Position))
	}
	b.WriteString("\n")
}

func renderRemoved(b *strings.Builder, blocks []model.RemovedBlock) {
	if len(blocks) == 0 {
		return
	}
	b.WriteString("## Removed\n\n")
	b.WriteString("| From | Destroy On Drop | Position |\n")
	b.WriteString("|------|-----------------|----------|\n")
	for _, r := range blocks {
		fmt.Fprintf(b, "| %s | %s | %s |\n",
			code(r.From), boolPtrCell(r.DestroyOnDrop), pos(r.Position))
	}
	b.WriteString("\n")
}

func renderChecks(b *strings.Builder, blocks []model.CheckBlock) {
	if len(blocks) == 0 {
		return
	}
	b.WriteString("## Checks\n\n")
	b.WriteString("| Name | Data Source | Assertions | Position |\n")
	b.WriteString("|------|-------------|:----------:|----------|\n")
	for _, c := range blocks {
		ds := "-"
		if c.DataSource != nil {
			ds = code(c.DataSource.Type + "." + c.DataSource.Name)
		}
		fmt.Fprintf(b, "| %s | %s | %d | %s |\n",
			code(c.Name), ds, len(c.Assertions), pos(c.Position))
	}
	b.WriteString("\n")
}

func renderSchemaFindings(b *strings.Builder, m *model.Module) {
	type row struct{ addr, kind, attr, message, rng string }
	var rows []row
	collect := func(addr string, f *model.SchemaFindings) {
		if f == nil {
			return
		}
		for _, a := range f.UnknownAttrs {
			rows = append(rows, row{addr, "unknown", a.Name, "", pos(a.Position)})
		}
		for _, d := range f.DeprecatedAttrs {
			rows = append(rows, row{addr, "deprecated", d.Name, oneLine(d.Message), pos(d.Position)})
		}
		for _, miss := range f.MissingRequired {
			rows = append(rows, row{addr, "missing required", miss, "", ""})
		}
	}
	for _, r := range m.ManagedResources {
		collect(r.Type+"."+r.Name, r.SchemaFindings)
	}
	for _, r := range m.EphemeralResources {
		collect(r.Type+"."+r.Name, r.SchemaFindings)
	}
	if len(rows) == 0 {
		return
	}
	b.WriteString("## Schema Findings\n\n")
	b.WriteString("| Resource | Kind | Attribute | Message | Position |\n")
	b.WriteString("|----------|------|-----------|---------|----------|\n")
	for _, r := range rows {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n",
			code(r.addr), r.kind, code(r.attr), r.message, r.rng)
	}
	b.WriteString("\n")
}

func renderDiagnostics(b *strings.Builder, diags model.Diagnostics) {
	if len(diags) == 0 {
		return
	}
	b.WriteString("## Diagnostics\n\n")
	b.WriteString("| Severity | Summary | Detail | Subject | Context |\n")
	b.WriteString("|----------|---------|--------|---------|---------|\n")
	for _, d := range diags {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n",
			string(d.Severity),
			oneLine(d.Summary),
			oneLine(d.Detail),
			posPtr(d.Subject),
			posPtr(d.Context))
	}
	b.WriteString("\n")
}

// --- cell helpers ---

// pos formats a source Position as `file.tf [L:C:B -> L:C:B]` (line:column:byte).
func pos(p model.Position) string {
	if p.Filename == "" {
		return ""
	}
	return fmt.Sprintf("`%s` [%d:%d:%d -> %d:%d:%d]",
		p.Filename,
		p.Start.Line, p.Start.Column, p.Start.Byte,
		p.End.Line, p.End.Column, p.End.Byte)
}

func posPtr(p *model.Position) string {
	if p == nil {
		return "-"
	}
	return pos(*p)
}

func exprCell(e *model.Expression) string {
	if e == nil {
		return "-"
	}
	return code(oneLine(e.Source))
}

func exprValueCell(e model.Expression) string {
	return code(oneLine(e.Source))
}

// typeCell renders a variable's type. It prefers the structured TypeSpec
// (open-inspector v0.7.0+), falling back to the raw typeexpr string.
func typeCell(raw string, spec *model.TypeSpec) string {
	if spec != nil {
		return code(typeSpecString(spec))
	}
	return code(raw)
}

// typeSpecString renders a structured TypeSpec tree as a single-line type
// expression, e.g. `object({name = string, tags = optional(map(string))})`.
func typeSpecString(t *model.TypeSpec) string {
	if t == nil {
		return ""
	}
	switch t.Kind {
	case model.TypeList, model.TypeSet, model.TypeMap:
		return string(t.Kind) + "(" + typeSpecString(t.Element) + ")"
	case model.TypeObject:
		names := make([]string, 0, len(t.Attributes))
		for name := range t.Attributes {
			names = append(names, name)
		}
		sort.Strings(names)
		parts := make([]string, 0, len(names))
		for _, name := range names {
			attr := t.Attributes[name]
			inner := typeSpecString(attr.Type)
			if attr.Optional {
				if attr.Default != nil {
					inner = "optional(" + inner + ", " + valueString(attr.Default) + ")"
				} else {
					inner = "optional(" + inner + ")"
				}
			}
			parts = append(parts, name+" = "+inner)
		}
		return "object({" + strings.Join(parts, ", ") + "})"
	case model.TypeTuple:
		parts := make([]string, 0, len(t.Elements))
		for _, el := range t.Elements {
			parts = append(parts, typeSpecString(el))
		}
		return "tuple([" + strings.Join(parts, ", ") + "])"
	case model.TypeDynamic:
		return "any"
	default:
		return string(t.Kind)
	}
}

// defaultCell renders a variable default. It prefers the decoded literal
// Value (open-inspector v0.7.0+), falling back to the raw expression source.
func defaultCell(expr *model.Expression, value *model.Value) string {
	if value != nil {
		return code(oneLine(valueString(value)))
	}
	return exprCell(expr)
}

// valueString renders a decoded constant Value as a single-line literal.
func valueString(v *model.Value) string {
	if v == nil {
		return ""
	}
	switch v.Kind {
	case model.ValueNull:
		return "null"
	case model.ValueString:
		return fmt.Sprintf("%q", v.String)
	case model.ValueNumber:
		return v.Number
	case model.ValueBool:
		return fmt.Sprintf("%t", v.Bool)
	case model.ValueList:
		parts := make([]string, 0, len(v.List))
		for _, el := range v.List {
			parts = append(parts, valueString(el))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case model.ValueTuple:
		parts := make([]string, 0, len(v.Tuple))
		for _, el := range v.Tuple {
			parts = append(parts, valueString(el))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case model.ValueMap:
		return objectString(v.Map)
	case model.ValueObject:
		return objectString(v.Object)
	default:
		return ""
	}
}

func objectString(entries map[string]*model.Value) string {
	names := make([]string, 0, len(entries))
	for name := range entries {
		names = append(names, name)
	}
	sort.Strings(names)
	parts := make([]string, 0, len(names))
	for _, name := range names {
		parts = append(parts, name+" = "+valueString(entries[name]))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

// refsCell renders the references an expression makes (open-inspector
// v0.6.0+) as a comma-separated list of canonical addresses.
func refsCell(refs []model.Reference) string {
	if len(refs) == 0 {
		return "-"
	}
	parts := make([]string, 0, len(refs))
	for _, r := range refs {
		addr := r.Address
		if r.Attribute != "" {
			addr += "." + r.Attribute
		}
		parts = append(parts, addr)
	}
	return code(strings.Join(parts, ", "))
}

// attrNamesCell lists the top-level attribute names captured in a resource's
// NestedBody (open-inspector v0.7.0+ replaces the old AttrNames slice).
func attrNamesCell(body *model.Body) string {
	if body == nil || len(body.Attributes) == 0 {
		return "-"
	}
	names := make([]string, 0, len(body.Attributes))
	for name := range body.Attributes {
		names = append(names, name)
	}
	sort.Strings(names)
	return code(strings.Join(names, ", "))
}

// inputsCell lists the input argument names passed to a module call.
func inputsCell(inputs map[string]model.Expression) string {
	if len(inputs) == 0 {
		return "-"
	}
	names := make([]string, 0, len(inputs))
	for name := range inputs {
		names = append(names, name)
	}
	sort.Strings(names)
	return code(strings.Join(names, ", "))
}

func boolCell(v bool) string {
	if v {
		return "✓"
	}
	return "-"
}

func boolPtrCell(v *bool) string {
	if v == nil {
		return "-"
	}
	if *v {
		return "✓"
	}
	return "✗"
}

func lifecycleCell(l *model.Lifecycle) string {
	if l == nil {
		return "-"
	}
	var parts []string
	if l.CreateBeforeDestroy != nil {
		parts = append(parts, fmt.Sprintf("create_before_destroy=%t", *l.CreateBeforeDestroy))
	}
	if l.PreventDestroy != nil {
		parts = append(parts, fmt.Sprintf("prevent_destroy=%t", *l.PreventDestroy))
	}
	if len(l.IgnoreChanges) > 0 {
		parts = append(parts, "ignore_changes=["+strings.Join(l.IgnoreChanges, ",")+"]")
	}
	if len(l.ReplaceTriggeredBy) > 0 {
		parts = append(parts, "replace_triggered_by=["+strings.Join(l.ReplaceTriggeredBy, ",")+"]")
	}
	if len(l.Preconditions) > 0 {
		parts = append(parts, fmt.Sprintf("preconditions=%d", len(l.Preconditions)))
	}
	if len(l.Postconditions) > 0 {
		parts = append(parts, fmt.Sprintf("postconditions=%d", len(l.Postconditions)))
	}
	if len(parts) == 0 {
		return "-"
	}
	return code(strings.Join(parts, "; "))
}

func mapCell(m map[string]string) string {
	if len(m) == 0 {
		return "-"
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var parts []string
	for _, k := range keys {
		parts = append(parts, k+"="+m[k])
	}
	return code(strings.Join(parts, ", "))
}

func code(s string) string {
	if s == "" {
		return "-"
	}
	return "`" + s + "`"
}

// oneLine collapses whitespace and escapes pipe characters so a value is safe
// to place inside a Markdown table cell.
func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	if s == "" {
		return "-"
	}
	return strings.ReplaceAll(s, "|", "\\|")
}

func sortedKeys(m map[string]model.ProviderRequirement) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
