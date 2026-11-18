package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderRequiredProviders(
	b *strings.Builder,
	cfg *config.RequiredConfig,
	reqs map[string]model.ProviderRequirement,
	pos func(model.Position) string,
) {
	if len(reqs) == 0 || !cfg.Render {
		// TODO: add logging output
		return
	}
	type reqRow struct {
		name string
		model.ProviderRequirement
	}
	rows := make([]reqRow, 0, len(reqs))
	for _, name := range sortedKeys(reqs) {
		rows = append(rows, reqRow{name, reqs[name]})
	}
	b.WriteString("## Required Providers\n\n")
	writeTable(b, []column[reqRow]{
		col("Name", cfg.Name, func(r reqRow) string { return code(r.name) }),
		col("Source", cfg.Source, func(r reqRow) string { return code(r.Source) }),
		col(
			"Version",
			cfg.Version,
			func(r reqRow) string { return code(strings.Join(r.VersionConstraints, ", ")) },
		),
		col(
			"Alias",
			cfg.Alias,
			func(r reqRow) string { return code(strings.Join(r.ConfigurationAliases, ", ")) },
		),
		col("Position", cfg.Position, func(r reqRow) string { return pos(r.Position) }),
	}, rows)
}

func renderProviders(
	b *strings.Builder,
	cfg *config.ProvidersConfig,
	providers []model.ProviderConfig,
	pos func(model.Position) string,
) {
	if len(providers) == 0 || !cfg.Render {
		// TODO: add logging output
		return
	}
	b.WriteString("## Providers\n\n")
	writeTable(b, []column[model.ProviderConfig]{
		col("Name", cfg.Name, func(p model.ProviderConfig) string {
			return code(p.Name)
		}),
		col("Alias", cfg.Alias, func(p model.ProviderConfig) string {
			return code(p.Alias)
		}),
		col("For Each", cfg.ForEach, func(p model.ProviderConfig) string {
			return exprCell(p.ForEach)
		}),
		col("Position", cfg.Position, func(p model.ProviderConfig) string {
			return pos(p.Position)
		}),
	}, providers)
}
