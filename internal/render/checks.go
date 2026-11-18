package render

import (
	"strconv"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderChecks(b *strings.Builder, cfg *config.ChecksConfig, blocks []model.CheckBlock, pos func(model.Position) string) {
	if len(blocks) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Checks\n\n")
	writeTable(b, []column[model.CheckBlock]{
		col("Name", cfg.Name, func(p model.CheckBlock) string {
			return code(p.Name)
		}),
		col("Data Source", cfg.DataSource, func(p model.CheckBlock) string {
			ds := "-"
			if p.DataSource != nil {
				return code(p.DataSource.Type + "." + p.DataSource.Name)
			}
			return ds
		}),
		col("Assertions", cfg.Assertions, func(p model.CheckBlock) string {
			return strconv.Itoa(len(p.Assertions))
		}),
		col("Position", cfg.Position, func(p model.CheckBlock) string {
			return pos(p.Position)
		}),
	}, blocks)
}
