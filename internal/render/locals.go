package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderLocals(
	b *strings.Builder,
	cfg *config.LocalsConfig,
	locals []model.Local,
	pos func(model.Position) string,
) {
	if len(locals) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Locals\n\n")
	writeTable(b, []column[model.Local]{
		col("Name", cfg.Name, func(p model.Local) string {
			return p.Name
		}),
		col("Value", cfg.Value, func(p model.Local) string {
			return exprValueCell(p.Value)
		}),
		col("References", cfg.References, func(p model.Local) string {
			return refsCell(p.Value.References)
		}),
		col("Position", cfg.Position, func(p model.Local) string {
			return pos(p.Position)
		}),
	}, locals)
}
