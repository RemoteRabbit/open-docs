package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderOutputs(
	b *strings.Builder,
	cfg *config.OutputsConfig,
	outs []model.Output,
	pos func(model.Position) string,
) {
	if len(outs) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Outputs\n\n")
	writeTable(b, []column[model.Output]{
		col("Name", cfg.Name, func(p model.Output) string {
			return p.Name
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
