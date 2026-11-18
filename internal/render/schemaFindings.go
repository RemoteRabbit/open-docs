package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderSchemaFindings(b *strings.Builder, cfg *config.SchemaConfig, m *model.Module, pos func(model.Position) string) {
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
			rows = append(
				rows,
				row{addr, "deprecated", d.Name, oneLine(d.Message), pos(d.Position)},
			)
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
	if len(rows) == 0 || !cfg.Render {
		return
	}

	b.WriteString("## Schema Findings\n\n")
	writeTable(b, []column[row]{
		col("Resource", cfg.Resource, func(p row) string {
			return code(p.addr)
		}),
		col("Kind", cfg.Kind, func(p row) string {
			return p.kind
		}),
		col("Attribute", cfg.Attribute, func(p row) string {
			return code(p.attr)
		}),
		col("Message", cfg.Message, func(p row) string {
			return p.message
		}),
		col("Position", cfg.Position, func(p row) string {
			return p.rng
		}),
	}, rows)
}
