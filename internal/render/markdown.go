// Package render turns an inspected *model.Module into documentation output.
//
// The current renderer is a verbose, debug-oriented view that dumps every
// field of every block type. It is meant for development and troubleshooting;
// the curated, reader-facing layout is a separate renderer (planned).
package render

import (
	"fmt"
	"path/filepath"
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
		fmt.Fprintf(
			&b,
			"<!-- VERSION-START -->\nThis documentation was generated with [open-docs v%s](https://github.com/RemoteRabbit/open-docs)\n<!-- VERSION-END -->\n\n",
			version,
		)
	}

	if len(m.RequiredCore) > 0 {
		fmt.Fprintf(&b, "- **Required core:** %s\n", code(strings.Join(m.RequiredCore, ", ")))
	}
	b.WriteString("\n")

	renderRequiredProviders(&b, cfg.Required, m.RequiredProviders, pos)
	renderProviders(&b, cfg.Providers, m.Providers, pos)
	renderVariables(&b, cfg.Variables, m.Variables, pos)
	renderOutputs(&b, cfg.Outputs, m.Outputs, pos)
	renderLocals(&b, cfg.Locals, m.Locals, pos)
	renderResources(&b, "Managed", cfg.Resources, m.ManagedResources, pos)
	renderResources(&b, "Data", cfg.Resources, m.DataResources, pos)
	renderResources(&b, "Ephemeral", cfg.Resources, m.EphemeralResources, pos)
	// renderModuleCalls(&b, m.ModuleCalls, pos)
	// renderMoved(&b, m.Moved, pos)
	// renderImports(&b, m.Imports, pos)
	// renderRemoved(&b, m.Removed, pos)
	// renderChecks(&b, m.Checks, pos)
	// renderSchemaFindings(&b, m, pos)
	renderDiagnostics(&b, cfg.Diagnostics, m.Diagnostics, pos)

	return b.String()
}

// func renderEphemeralResources(
// 	b *strings.Builder,
// 	resources []model.EphemeralResource,
// 	pos func(model.Position) string,
// ) {
// 	if len(resources) == 0 {
// 		return
// 	}
// 	b.WriteString("## Ephemeral Resources\n\n")
// 	b.WriteString(
// 		"| Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |\n",
// 	)
// 	b.WriteString(
// 		"|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|\n",
// 	)
// 	for _, r := range resources {
// 		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
// 			code(r.Type),
// 			code(r.Name),
// 			code(r.Provider),
// 			exprCell(r.Count),
// 			exprCell(r.ForEach),
// 			code(strings.Join(r.DependsOn, ", ")),
// 			attrNamesCell(r.NestedBody),
// 			lifecycleCell(r.Lifecycle),
// 			oneLine(r.Comment),
// 			pos(r.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderModuleCalls(
// 	b *strings.Builder,
// 	calls []model.ModuleCall,
// 	pos func(model.Position) string,
// ) {
// 	if len(calls) == 0 {
// 		return
// 	}
// 	b.WriteString("## Module Calls\n\n")
// 	b.WriteString(
// 		"| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |\n",
// 	)
// 	b.WriteString(
// 		"|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|\n",
// 	)
// 	for _, c := range calls {
// 		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
// 			code(c.Name),
// 			code(c.Source),
// 			code(c.Version),
// 			exprCell(c.Count),
// 			exprCell(c.ForEach),
// 			code(strings.Join(c.DependsOn, ", ")),
// 			mapCell(c.Providers),
// 			inputsCell(c.Inputs),
// 			oneLine(c.Comment),
// 			pos(c.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderMoved(b *strings.Builder, blocks []model.MovedBlock, pos func(model.Position) string) {
// 	if len(blocks) == 0 {
// 		return
// 	}
// 	b.WriteString("## Moved\n\n")
// 	b.WriteString("| From | To | Position |\n")
// 	b.WriteString("|------|----|----------|\n")
// 	for _, m := range blocks {
// 		fmt.Fprintf(b, "| %s | %s | %s |\n", code(m.From), code(m.To), pos(m.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderImports(
// 	b *strings.Builder,
// 	blocks []model.ImportBlock,
// 	pos func(model.Position) string,
// ) {
// 	if len(blocks) == 0 {
// 		return
// 	}
// 	b.WriteString("## Imports\n\n")
// 	b.WriteString("| To | ID | Provider | Position |\n")
// 	b.WriteString("|----|----|----------|----------|\n")
// 	for _, im := range blocks {
// 		fmt.Fprintf(b, "| %s | %s | %s | %s |\n",
// 			code(im.To), exprValueCell(im.ID), code(im.Provider), pos(im.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderRemoved(
// 	b *strings.Builder,
// 	blocks []model.RemovedBlock,
// 	pos func(model.Position) string,
// ) {
// 	if len(blocks) == 0 {
// 		return
// 	}
// 	b.WriteString("## Removed\n\n")
// 	b.WriteString("| From | Destroy On Drop | Position |\n")
// 	b.WriteString("|------|-----------------|----------|\n")
// 	for _, r := range blocks {
// 		fmt.Fprintf(b, "| %s | %s | %s |\n",
// 			code(r.From), boolPtrCell(r.DestroyOnDrop), pos(r.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderChecks(b *strings.Builder, blocks []model.CheckBlock, pos func(model.Position) string) {
// 	if len(blocks) == 0 {
// 		return
// 	}
// 	b.WriteString("## Checks\n\n")
// 	b.WriteString("| Name | Data Source | Assertions | Position |\n")
// 	b.WriteString("|------|-------------|:----------:|----------|\n")
// 	for _, c := range blocks {
// 		ds := "-"
// 		if c.DataSource != nil {
// 			ds = code(c.DataSource.Type + "." + c.DataSource.Name)
// 		}
// 		fmt.Fprintf(b, "| %s | %s | %d | %s |\n",
// 			code(c.Name), ds, len(c.Assertions), pos(c.Position))
// 	}
// 	b.WriteString("\n")
// }
//
// func renderSchemaFindings(b *strings.Builder, m *model.Module, pos func(model.Position) string) {
// 	type row struct{ addr, kind, attr, message, rng string }
// 	var rows []row
// 	collect := func(addr string, f *model.SchemaFindings) {
// 		if f == nil {
// 			return
// 		}
// 		for _, a := range f.UnknownAttrs {
// 			rows = append(rows, row{addr, "unknown", a.Name, "", pos(a.Position)})
// 		}
// 		for _, d := range f.DeprecatedAttrs {
// 			rows = append(
// 				rows,
// 				row{addr, "deprecated", d.Name, oneLine(d.Message), pos(d.Position)},
// 			)
// 		}
// 		for _, miss := range f.MissingRequired {
// 			rows = append(rows, row{addr, "missing required", miss, "", ""})
// 		}
// 	}
// 	for _, r := range m.ManagedResources {
// 		collect(r.Type+"."+r.Name, r.SchemaFindings)
// 	}
// 	for _, r := range m.EphemeralResources {
// 		collect(r.Type+"."+r.Name, r.SchemaFindings)
// 	}
// 	if len(rows) == 0 {
// 		return
// 	}
// 	b.WriteString("## Schema Findings\n\n")
// 	b.WriteString("| Resource | Kind | Attribute | Message | Position |\n")
// 	b.WriteString("|----------|------|-----------|---------|----------|\n")
// 	for _, r := range rows {
// 		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n",
// 			code(r.addr), r.kind, code(r.attr), r.message, r.rng)
// 	}
// 	b.WriteString("\n")
// }
