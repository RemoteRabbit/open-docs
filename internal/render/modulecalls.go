package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderModuleCalls(
	b *strings.Builder,
	cfg *config.ModulesConfig,
	calls []model.ModuleCall,
	pos func(model.Position) string,
) {
	if len(calls) == 0 || !cfg.Render {
		return
	}
	b.WriteString("## Module Calls\n\n")
	writeTable(b, []column[model.ModuleCall]{
		col("Name", cfg.Name, func(p model.ModuleCall) string {
			return string(p.Name)
		}),
		col("Source", cfg.Source, func(p model.ModuleCall) string {
			return string(p.Source)
		}),
		col("Version", cfg.Version, func(p model.ModuleCall) string {
			return string(p.Version)
		}),
		col("Count", cfg.Count, func(p model.ModuleCall) string {
			return exprCell(p.Count)
		}),
		col("For Each", cfg.ForEach, func(p model.ModuleCall) string {
			return exprCell(p.ForEach)
		}),
		col("Depends On", cfg.DependsOn, func(p model.ModuleCall) string {
			return code(strings.Join(p.DependsOn, ", "))
		}),
		col("Providers", cfg.Providers, func(p model.ModuleCall) string {
			return mapCell(p.Providers)
		}),
		col("Inputs", cfg.Inputs, func(p model.ModuleCall) string {
			return inputsCell(p.Inputs)
		}),
		col("Comment", cfg.Comment, func(p model.ModuleCall) string {
			return oneLine(p.Comment)
		}),
		col("Position", cfg.Position, func(p model.ModuleCall) string {
			return pos(p.Position)
		}),
	}, calls)
}
