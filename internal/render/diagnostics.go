package render

import (
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

func renderDiagnostics(
	b *strings.Builder,
	cfg *config.DiagnosticsConfig,
	diags []model.Diagnostic,
	pos func(model.Position) string,
) {
	if len(diags) == 0 || !cfg.Render {
		return
	}

	b.WriteString("## Diagnostics\n\n")
	writeTable(b, []column[model.Diagnostic]{
		col("Severity", cfg.Severity, func(p model.Diagnostic) string {
			return string(p.Severity)
		}),
		col("Summary", cfg.Summary, func(p model.Diagnostic) string {
			return string(p.Summary)
		}),
		col("Detail", cfg.Detail, func(p model.Diagnostic) string {
			return string(p.Detail)
		}),
		col("Subject", cfg.Subject, func(p model.Diagnostic) string {
			return posPtr(p.Subject, pos)
		}),
		col("Context", cfg.Context, func(p model.Diagnostic) string {
			return posPtr(p.Context, pos)
		}),
	}, diags)
}
