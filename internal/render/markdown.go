// Package render turns an inspected *model.Module into documentation output.
//
// The current renderer is a verbose, debug-oriented view that dumps every
// field of every block type. It is meant for development and troubleshooting;
// the curated, reader-facing layout is a separate renderer (planned).
package render

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-inspector/pkg/model"
)

type column[T any] struct {
	header  string
	align   string
	enabled bool
	cell    func(T) string
}

// Markdown renders a module as a verbose Markdown reference document. Source
// links are made relative to outputFile, or to the module directory when the
// document is written to stdout.
func Markdown(m *model.Module, outputFile string, cfg *config.Config, version string) string {
	var b strings.Builder
	linkBase := m.Path
	if outputFile != "" {
		if absoluteOutput, err := filepath.Abs(outputFile); err == nil {
			linkBase = filepath.Dir(absoluteOutput)
		}
	}

	pos := func(p model.Position) string {
		return position(Position{Position: p, LinkBase: linkBase, Verbose: cfg.Main.Verbose})
	}

	if cfg.Main.Header {
		fmt.Fprintf(&b, "# Module `%s`\n\n", filepath.Base(m.Path))
	}

	// Add version information
	if cfg.Main.Version {
		if version != "" {
			fmt.Fprintf(
				&b,
				"<!-- VERSION-START -->\nThis documentation was generated with [open-docs %s](https://github.com/RemoteRabbit/open-docs)\n<!-- VERSION-END -->\n\n",
				version,
			)
		}
	}

	if len(m.RequiredCore) > 0 {
		fmt.Fprintf(&b, "- **Required core:** %s\n", code(strings.Join(m.RequiredCore, ", ")))
	}
	b.WriteString("\n")

	renderRequiredProviders(&b, cfg.Required, m.RequiredProviders, pos)
	renderProviders(&b, cfg.Providers, m.Providers, pos)
	renderVariables(&b, cfg.Variables, m.Variables, pos)
	renderOutputs(&b, cfg.Outputs, m.Outputs, pos)
	renderLocals(&b, cfg.Locals, m.Locals, pos)
	renderResources(&b, "Managed", cfg.Resources, m.ManagedResources, pos)
	renderResources(&b, "Data", cfg.Data, m.DataResources, pos)
	renderResources(&b, "Ephemeral", cfg.Ephemeral, m.EphemeralResources, pos)
	renderModuleCalls(&b, cfg.Modules, m.ModuleCalls, pos)
	renderMoved(&b, cfg.Moved, m.Moved, pos)
	renderImports(&b, cfg.Imports, m.Imports, pos)
	renderRemoved(&b, cfg.Removed, m.Removed, pos)
	renderChecks(&b, cfg.Checks, m.Checks, pos)
	renderSchemaFindings(&b, cfg.Schema, m, pos)
	renderDiagnostics(&b, cfg.Diagnostics, m.Diagnostics, pos)

	// Each section ends with a blank line; end the document with exactly one
	// newline so the file is stable under end-of-file fixers.
	return strings.TrimRight(b.String(), "\n") + "\n"
}
