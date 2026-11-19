// Command open-docs generates documentation for a Terraform/OpenTofu module by
// inspecting it with open-inspector.
//
// Usage:
//
//	open-docs [flags] <module-dir>
//
// Flags:
//
//	-o <file>       Write output to a file (replace mode); overrides config.
//	-config <file>  Use this config file instead of discovering .open-docs.hcl.
//	-schema         Enrich with provider schema (shells out to tofu/terraform).
//
// Output goes to stdout by default. Set the destination and mode in
// .open-docs.hcl with main { docsfile, mode }. The -o flag is a shortcut for
// replace mode to the given file and takes precedence over config.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/remoterabbit/open-docs/internal/config"
	"github.com/remoterabbit/open-docs/internal/output"
	"github.com/remoterabbit/open-docs/internal/render"
	"github.com/remoterabbit/open-inspector/pkg/inspector"
)

// version is set at build time with -ldflags "-X main.version=..." (GoReleaser
// does this for releases, and the Makefile sets "dev").
var version string

func main() {
	outFlag := flag.String("o", "", "write output to this file in replace mode (overrides config)")
	cfgFlag := flag.String(
		"config",
		"",
		"path to config file (default: discover "+config.DefaultFilename+")",
	)
	schemaFlag := flag.Bool("schema", false, "enrich with provider schema via tofu/terraform")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: open-docs [flags] <module-dir>\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	dir := flag.Arg(0)

	if err := run(dir, *outFlag, *cfgFlag, *schemaFlag); err != nil {
		fmt.Fprintf(os.Stderr, "open-docs: %v\n", err)
		os.Exit(1)
	}
}

// run [TODO:description]
func run(dir, outFlag, configFlag string, useSchema bool) error {
	cfg, err := loadConfig(dir, configFlag)
	if err != nil {
		return err
	}
	file, mode := resolveOutput(cfg, dir, outFlag)

	var opts []inspector.Option
	if useSchema {
		opts = append(opts, inspector.WithSchemaAuto())
	}
	mod, err := inspector.Inspect(dir, opts...)
	if err != nil {
		return err
	}

	doc := render.Markdown(mod, file, cfg, displayVersion())

	if err := output.Write(os.Stdout, file, mode, doc); err != nil {
		return err
	}
	if mode != config.ModeStdout && file != "" {
		fmt.Fprintf(os.Stderr, "wrote %s (%s)\n", file, mode)
	}
	return nil
}

// loadConfig loads the explicit config file, or discovers .open-docs.hcl near
// the module, falling back to defaults when none exists.
func loadConfig(dir, configFlag string) (*config.Config, error) {
	path := configFlag
	if path == "" {
		if found, ok := config.Find(dir); ok {
			path = found
		}
	}
	if path == "" {
		return config.Default(), nil
	}
	return config.Load(path)
}

// resolveOutput merges config output settings with the -o flag. The flag, when
// set, wins: it forces replace mode to that file (relative to the working
// directory). Config file paths resolve relative to the module directory.
func resolveOutput(cfg *config.Config, dir, outFlag string) (file, mode string) {
	mode = config.ModeStdout
	if cfg.Main != nil && cfg.Main.DocsFile != "" {
		file = cfg.Main.DocsFile
		if !filepath.IsAbs(file) {
			file = filepath.Join(dir, file)
		}

		mode = cfg.Main.Mode
		if mode == "" {
			mode = config.ModeReplace
		}
	}
	if outFlag != "" {
		file = outFlag
		mode = config.ModeReplace
	}
	return file, mode
}

// displayVersion returns the version shown in generated docs: the build-time
// version, else the module version recorded by `go install module@version`,
// else "dev".
func displayVersion() string {
	if version == "dev" {
		return version
	}
	if version != "" {
		return "v" + strings.TrimPrefix(version, "v")
	}
	if bi, ok := debug.ReadBuildInfo(); ok && strings.HasPrefix(bi.Main.Version, "v") &&
		!strings.Contains(bi.Main.Version, "+dirty") {
		return bi.Main.Version
	}
	return "dev"
}
