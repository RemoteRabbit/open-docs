package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderRemoved(
	b *strings.Builder,
	cfg *config.RemovedConfig,
	blocks []model.RemovedBlock,
	pos func(model.Position) string,
) {
	if len(blocks) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Removed\n\n")
	writeTable(b, []column[model.RemovedBlock]{
		col("From", cfg.From, func(p model.RemovedBlock) string {
			return code(p.From)
		}),
		col("Destroy on Drop", cfg.DestroyOnDrop, func(p model.RemovedBlock) string {
			return boolPtrCell(p.DestroyOnDrop)
		}),
		col("Position", cfg.Position, func(p model.RemovedBlock) string {
			return pos(p.Position)
		}),
	}, blocks)
}
