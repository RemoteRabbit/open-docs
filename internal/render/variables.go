package render

import (
	"fmt"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderVariables(
	b *strings.Builder,
	cfg *config.VariablesConfig,
	vars []model.Variable,
	pos func(model.Position) string,
) {
	if len(vars) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Inputs\n\n")
	writeTable(b, []column[model.Variable]{
		col("Name", cfg.Name, func(p model.Variable) string {
			return code(p.Name)
		}),
		col("Type", cfg.Type, func(p model.Variable) string {
			return typeCell(p.Type, p.TypeSpec)
		}),
		col("Default", cfg.Default, func(p model.Variable) string {
			return defaultCell(p.Default, p.DefaultValue)
		}),
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
		col("Position", cfg.Position, func(p model.Variable) string {
			return pos(p.Position)
		}),
	}, vars)
}
