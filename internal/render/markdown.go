// Package render turns an inspected *model.Module into documentation output.
//
// The current renderer is a verbose, debug-oriented view that dumps every
// field of every block type. It is meant for development and troubleshooting;
// the curated, reader-facing layout is a separate renderer (planned). See
// ROADMAP.md.
package render

import (
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

type column[T any] struct {
	header  string
	align   string
	enabled bool
	cell    func(T) string
}

// col builds a left-aligned column; chain .center() for other alignment.
func col[T any](header string, enabled bool, cell func(T) string) column[T] {
	return column[T]{header: header, enabled: enabled, cell: cell}
}

func (c column[T]) center() column[T] { c.align = "center"; return c }

// sepFor returns a delimiter cell as wide as the rendered header cell
// (" Name " = len+2), with optional alignment colons.
func sepFor(header, align string) string {
	n := len(header) + 2
	switch align {
	case "center":
		return ":" + strings.Repeat("-", n-2) + ":"
	case "right":
		return strings.Repeat("-", n-1) + ":"
	default:
		return strings.Repeat("-", n)
	}
}

func writeTable[T any](b *strings.Builder, cols []column[T], rows []T) {
	var active []column[T]
	for _, c := range cols {
		if c.enabled {
			active = append(active, c)
		}
	}
	if len(active) == 0 {
		return
	}
	headers := make([]string, len(active))
	seps := make([]string, len(active))
	for i, c := range active {
		headers[i], seps[i] = c.header, sepFor(c.header, c.align)
	}
	fmt.Fprintf(b, "| %s |\n|%s|\n", strings.Join(headers, " | "), strings.Join(seps, "|"))
	for _, r := range rows {
		cells := make([]string, len(active))
		for i, c := range active {
			cells[i] = c.cell(r)
		}
		fmt.Fprintf(b, "| %s |\n", strings.Join(cells, " | "))
	}
	b.WriteString("\n")
}

// Markdown renders a module as a verbose Markdown reference document. Source
// links are made relative to outputFile, or to the module directory when the
// document is written to stdout.
func Markdown(m *model.Module, outputFile string, cfg *config.Config, version string) string {
	var b strings.Builder
	linkBase := m.Path
	if outputFile != "" {
		if absoluteOutput, err := filepath.Abs(outputFile); err == nil {
			linkBase = filepath.Dir(absoluteOutput)
		}
	}
	pos := func(p model.Position) string {
		return position(p, linkBase)
	}
	if cfg.Main.Header {
		fmt.Fprintf(&b, "# Module `%s`\n\n", filepath.Base(m.Path))
	}

	// Add version information
	if version != "" {
		fmt.Fprintf(&b, "<!-- VERSION-START -->\nThis documentation was generated with [open-docs v%s](https://github.com/RemoteRabbit/open-docs)\n<!-- VERSION-END -->\n\n", version)
	}

	if len(m.RequiredCore) > 0 {
		fmt.Fprintf(&b, "- **Required core:** %s\n", code(strings.Join(m.RequiredCore, ", ")))
	}
	b.WriteString("\n")

	renderRequiredProviders(&b, cfg.Required, m.RequiredProviders, pos)
	renderProviders(&b, cfg.Providers, m.Providers, pos)
	renderVariables(&b, cfg.Variables, m.Variables, pos)
	renderOutputs(&b, cfg.Outputs, m.Outputs, pos)
	renderLocals(&b, m.Locals, pos)
	renderResources(&b, "Managed Resources", m.ManagedResources, pos)
	renderResources(&b, "Data Resources", m.DataResources, pos)
	renderEphemeralResources(&b, m.EphemeralResources, pos)
	renderModuleCalls(&b, m.ModuleCalls, pos)
	renderMoved(&b, m.Moved, pos)
	renderImports(&b, m.Imports, pos)
	renderRemoved(&b, m.Removed, pos)
	renderChecks(&b, m.Checks, pos)
	renderSchemaFindings(&b, m, pos)
	renderDiagnostics(&b, m.Diagnostics, pos)

	return b.String()
}

func renderRequiredProviders(b *strings.Builder, cfg *config.RequiredConfig, reqs map[string]model.ProviderRequirement, pos func(model.Position) string) {
	if len(reqs) == 0 || !cfg.Render {
		// TODO: add logging output
		return
	}
	type reqRow struct {
		name string
		model.ProviderRequirement
	}
	rows := make([]reqRow, 0, len(reqs))
	for _, name := range sortedKeys(reqs) {
		rows = append(rows, reqRow{name, reqs[name]})
	}
	b.WriteString("## Required Providers\n\n")
	writeTable(b, []column[reqRow]{
		col("Name", cfg.Name, func(r reqRow) string { return code(r.name) }),
		col("Source", cfg.Source, func(r reqRow) string { return code(r.Source) }),
		col("Version", cfg.Version, func(r reqRow) string { return code(strings.Join(r.VersionConstraints, ", ")) }),
		col("Alias", cfg.Alias, func(r reqRow) string { return code(strings.Join(r.ConfigurationAliases, ", ")) }),
		col("Position", cfg.Position, func(r reqRow) string { return pos(r.Position) }),
	}, rows)
}

func renderProviders(b *strings.Builder, cfg *config.ProvidersConfig, providers []model.ProviderConfig, pos func(model.Position) string) {
	if len(providers) == 0 || !cfg.Render {
		// TODO: add logging output
		return
	}
	b.WriteString("## Providers\n\n")
	writeTable(b, []column[model.ProviderConfig]{
		col("Name", cfg.Name, func(p model.ProviderConfig) string { return code(p.Name) }),
		col("Alias", cfg.Alias, func(p model.ProviderConfig) string { return code(p.Alias) }),
		col("For Each", cfg.ForEach, func(p model.ProviderConfig) string { return exprCell(p.ForEach) }),
		col("Position", cfg.Position, func(p model.ProviderConfig) string { return pos(p.Position) }),
	}, providers)
}

func renderVariables(b *strings.Builder, cfg *config.VariablesConfig, vars []model.Variable, pos func(model.Position) string) {
	if len(vars) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Inputs\n\n")
	writeTable(b, []column[model.Variable]{
		col("Name", cfg.Name, func(p model.Variable) string {
			return code(p.Name)
		}),
		col(
			"Type",
			cfg.Type,
			func(p model.Variable) string { return typeCell(p.Type, p.TypeSpec) },
		),
		col(
			"Default",
			cfg.Default,
			func(p model.Variable) string { return defaultCell(p.Default, p.DefaultValue) },
		),
		col("Description", cfg.Description, func(p model.Variable) string {
			return p.Description
		}),
		// TODO: should make this configurable if there are comments or not as well
		col("Comment", cfg.Comment, func(p model.Variable) string {
			return p.Comment
		}),
		col("Sensitive", cfg.Sensitive, func(p model.Variable) string {
			return boolCell(p.Sensitive)
		}),
		col("Nullable", cfg.Nullable, func(p model.Variable) string {
			return boolPtrCell(p.Nullable)
		}),
		col("Ephemeral", cfg.Ephemeral, func(p model.Variable) string {
			return boolCell(p.Ephemeral)
		}),
		// TODO: render out some or all validations, thinking on how this should be formatted
		col("Validation Count", cfg.Validation, func(p model.Variable) string {
			return fmt.Sprintf("%d", len(p.Validations))
		}),
		// {"Validation"},
		col("Position", cfg.Position, func(p model.Variable) string {
			return pos(p.Position)
		}),
	}, vars)
}

func renderOutputs(b *strings.Builder, cfg *config.OutputsConfig, outs []model.Output, pos func(model.Position) string) {
	if len(outs) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Outputs\n\n")
	writeTable(b, []column[model.Output]{
		col("Name", cfg.Name, func(p model.Output) string {
			return code(p.Name)
		}),
		col("Value", cfg.Value, func(p model.Output) string {
			return exprValueCell(p.Value)
		}),
		col("References", cfg.References, func(p model.Output) string {
			return refsCell(p.Value.References)
		}),
		col("Description", cfg.Description, func(p model.Output) string {
			return p.Description
		}),
		col("Comment", cfg.Comment, func(p model.Output) string {
			return p.Comment
		}),
		col("Sensitive", cfg.Sensitive, func(p model.Output) string {
			return boolCell(p.Sensitive)
		}),
		col("Ephemeral", cfg.Ephemeral, func(p model.Output) string {
			return boolCell(p.Ephemeral)
		}),
		col("Depends On", cfg.DependsOn, func(p model.Output) string {
			return code(strings.Join(p.DependsOn, ", "))
		}),
		col("Position", cfg.Position, func(p model.Output) string {
			return pos(p.Position)
		}),
	}, outs)
}

func renderLocals(b *strings.Builder, locals []model.Local, pos func(model.Position) string) {
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

func renderResources(b *strings.Builder, heading string, resources []model.Resource, pos func(model.Position) string) {
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

func renderEphemeralResources(b *strings.Builder, resources []model.EphemeralResource, pos func(model.Position) string) {
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

func renderModuleCalls(b *strings.Builder, calls []model.ModuleCall, pos func(model.Position) string) {
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

func renderMoved(b *strings.Builder, blocks []model.MovedBlock, pos func(model.Position) string) {
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

func renderImports(b *strings.Builder, blocks []model.ImportBlock, pos func(model.Position) string) {
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

func renderRemoved(b *strings.Builder, blocks []model.RemovedBlock, pos func(model.Position) string) {
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

func renderChecks(b *strings.Builder, blocks []model.CheckBlock, pos func(model.Position) string) {
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

func renderSchemaFindings(b *strings.Builder, m *model.Module, pos func(model.Position) string) {
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

func renderDiagnostics(b *strings.Builder, diags model.Diagnostics, pos func(model.Position) string) {
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
			posPtr(d.Subject, pos),
			posPtr(d.Context, pos))
	}
	b.WriteString("\n")
}

// --- cell helpers ---

// position formats a source Position as a document-relative source link followed
// by its line:column:byte range.
func position(p model.Position, linkBase string) string {
	if p.Filename == "" {
		return ""
	}
	target, err := filepath.Rel(linkBase, filepath.FromSlash(p.Filename))
	if err != nil {
		target = p.Filename
	} else if !strings.HasPrefix(target, ".") {
		target = "./" + target
	}
	target = filepath.ToSlash(target)
	href := (&url.URL{Path: target}).EscapedPath()
	if p.Start.Line > 0 {
		href += fmt.Sprintf("#L%d", p.Start.Line)
		if p.End.Line > p.Start.Line {
			href += fmt.Sprintf("-L%d", p.End.Line)
		}
	}
	return fmt.Sprintf("[`%s`](%s) [%d:%d:%d -> %d:%d:%d]",
		target, href,
		p.Start.Line, p.Start.Column, p.Start.Byte,
		p.End.Line, p.End.Column, p.End.Byte)
}

func posPtr(p *model.Position, pos func(model.Position) string) string {
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
