package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderImports(
	b *strings.Builder,
	cfg *config.ImportsConfig,
	blocks []model.ImportBlock,
	pos func(model.Position) string,
) {
	if len(blocks) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Imports\n\n")
	writeTable(b, []column[model.ImportBlock]{
		col("To", cfg.To, func(p model.ImportBlock) string {
			return code(p.To)
		}),
		col("ID", cfg.Id, func(p model.ImportBlock) string {
			return exprValueCell(p.ID)
		}),
		col("Provider", cfg.Provider, func(p model.ImportBlock) string {
			return code(p.Provider)
		}),
		col("Position", cfg.Position, func(p model.ImportBlock) string {
			return pos(p.Position)
		}),
	}, blocks)
}
