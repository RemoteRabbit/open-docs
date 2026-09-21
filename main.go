// Command open-doc generates documentation for a Terraform/OpenTofu module by
// inspecting it with open-inspector.
//
// Usage:
//
//	open-doc [flags] <module-dir>
//
// Flags:
//
//	-o <file>       Write output to a file (replace mode); overrides config.
//	-config <file>  Use this config file instead of discovering .open-doc.hcl.
//	-schema         Enrich with provider schema (shells out to tofu/terraform).
//
// Output destination and mode default to stdout, and can be set in
// .open-doc.hcl via an output { file, mode } block. The -o flag is a shortcut
// for replace mode to the given file and takes precedence over config.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/remoterabbit/open-doc/internal/config"
	"github.com/remoterabbit/open-doc/internal/output"
	"github.com/remoterabbit/open-doc/internal/render"
	"github.com/remoterabbit/open-inspector/pkg/inspector"
)

func main() {
	outFlag := flag.String("o", "", "write output to this file in replace mode (overrides config)")
	configFlag := flag.String("config", "", "path to config file (default: discover "+config.DefaultFilename+")")
	useSchema := flag.Bool("schema", false, "enrich with provider schema via tofu/terraform")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: open-doc [flags] <module-dir>\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	dir := flag.Arg(0)

	if err := run(dir, *outFlag, *configFlag, *useSchema); err != nil {
		fmt.Fprintf(os.Stderr, "open-doc: %v\n", err)
		os.Exit(1)
	}
}

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

	doc := render.Markdown(mod)

	if err := output.Write(os.Stdout, file, mode, doc); err != nil {
		return err
	}
	if mode != config.ModeStdout && file != "" {
		fmt.Fprintf(os.Stderr, "wrote %s (%s)\n", file, mode)
	}
	return nil
}

// loadConfig loads the explicit config file, or discovers .open-doc.hcl near
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
	if cfg.Output != nil && cfg.Output.File != "" {
		file = cfg.Output.File
		if !filepath.IsAbs(file) {
			file = filepath.Join(dir, file)
		}
		mode = cfg.Output.Mode
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
