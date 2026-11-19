package render

import (
	"fmt"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderResources(
	b *strings.Builder,
	title string,
	cfg *config.ResourcesConfig,
	resources []model.Resource,
	pos func(model.Position) string,
) {
	if len(resources) == 0 || !cfg.Render {
		return
	}
	fmt.Fprintf(b, "## %s Resources\n\n", title)
	writeTable(b, []column[model.Resource]{
		col("Name", cfg.Name, func(p model.Resource) string {
			return p.Name
		}),
		col("Mode", cfg.Mode, func(p model.Resource) string {
			return string(p.Mode)
		}),
		col("Type", cfg.Type, func(p model.Resource) string {
			return code(p.Type)
		}),
		col("Provider", cfg.Provider, func(p model.Resource) string {
			return p.Provider
		}),
		// TODO: look into having a merged column for count and for_each since you would only have one
		// or the other, but also make each toggle.
		col("Count", cfg.Count, func(p model.Resource) string {
			return exprCell(p.Count)
		}),
		col("For Each", cfg.ForEach, func(p model.Resource) string {
			return exprCell(p.ForEach)
		}),
		col("Depends On", cfg.DependsOn, func(p model.Resource) string {
			return strings.Join(p.DependsOn, ", ")
		}),
		col("Nested Body", cfg.NestedBody, func(p model.Resource) string {
			return attrNamesCell(p.NestedBody)
		}),
		col("LifeCycle", cfg.Lifecycle, func(p model.Resource) string {
			return lifecycleCell(p.Lifecycle)
		}),
		col("Comment", cfg.Comment, func(p model.Resource) string {
			return oneLine(p.Comment)
		}),
		col("Position", cfg.Position, func(p model.Resource) string {
			return pos(p.Position)
		}),
	}, resources)
}
