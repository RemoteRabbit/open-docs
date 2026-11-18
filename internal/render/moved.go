package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderMoved(b *strings.Builder, cfg *config.MovedConfig, blocks []model.MovedBlock, pos func(model.Position) string) {
	if len(blocks) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Moved\n\n")
	writeTable(b, []column[model.MovedBlock]{
		col("From", cfg.From, func(p model.MovedBlock) string {
			return oneLine(p.From)
		}),
		col("To", cfg.To, func(p model.MovedBlock) string {
			return oneLine(p.To)
		}),
		col("Position", cfg.Position, func(p model.MovedBlock) string {
			return pos(p.Position)
		}),
	}, blocks)
}
