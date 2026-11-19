// Command open-docs generates documentation for a Terraform/OpenTofu module by
// inspecting it with open-inspector.
//
// Usage:
//
//	open-docs [flags] <module-dir>
//
// Flags:
//
//	-m <mode>       Output mode: stdout, replace, or inject; overrides config.
//	-o <file>       Write output to a file (replace mode unless -m is set);
//	                overrides config.
//	-config <file>  Use this config file instead of discovering .open-docs.hcl.
//	-show <list>    Render these comma-separated sections; overrides config.
//	-hide <list>    Skip these comma-separated sections; overrides config and -show.
//	-schema         Enrich with provider schema (shells out to tofu/terraform).
//	-check          Exit non-zero if the docs file is out of date; writes nothing.
//	-version        Print the open-docs version and exit.
//
// Output goes to stdout by default. Set the destination and mode in
// .open-docs.hcl with main { docsfile, mode }. The -m and -o flags take
// precedence over config; -o alone is a shortcut for replace mode to the given
// file.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
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

// flags holds the parsed command-line flags.
type flags struct {
	mode, out, config, show, hide string
	schema, check                 bool
}

func main() {
	var f flags
	flag.StringVar(&f.mode, "m", "", "output `mode`: stdout, replace, or inject (overrides config)")
	flag.StringVar(
		&f.out,
		"o",
		"",
		"write output to this `file`, in replace mode unless -m is set (overrides config)",
	)
	flag.StringVar(
		&f.config,
		"config",
		"",
		"path to config `file` (default: discover "+config.DefaultFilename+")",
	)
	flag.StringVar(&f.show, "show", "", "comma-separated `sections` to render (overrides config)")
	flag.StringVar(
		&f.hide,
		"hide",
		"",
		"comma-separated `sections` to skip (overrides config and -show)",
	)
	flag.BoolVar(&f.schema, "schema", false, "enrich with provider schema via tofu/terraform")
	flag.BoolVar(
		&f.check,
		"check",
		false,
		"exit non-zero if the docs file is out of date, without writing it",
	)
	showVersion := flag.Bool("version", false, "print the open-docs version and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: open-docs [flags] <module-dir>\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Println("open-docs", displayVersion())
		return
	}
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(flag.Arg(0), f); err != nil {
		fmt.Fprintf(os.Stderr, "open-docs: %v\n", err)
		os.Exit(1)
	}
}

// run renders the docs for the module in dir and writes them where the flags
// and config say, or with -check only compares them to the docs file.
func run(dir string, f flags) error {
	switch f.mode {
	case "", config.ModeStdout, config.ModeReplace, config.ModeInject:
	default:
		return fmt.Errorf("-m %q is invalid (want %q, %q, or %q)",
			f.mode, config.ModeStdout, config.ModeReplace, config.ModeInject)
	}

	cfg, err := loadConfig(dir, f.config)
	if err != nil {
		return err
	}
	if err := setRender(cfg, f.show, true); err != nil {
		return fmt.Errorf("-show: %w", err)
	}
	if err := setRender(cfg, f.hide, false); err != nil {
		return fmt.Errorf("-hide: %w", err)
	}
	file, mode := resolveOutput(cfg, dir, f.mode, f.out)
	if f.check && (mode == config.ModeStdout || file == "") {
		return fmt.Errorf("-check needs replace or inject mode and a docs file")
	}

	var opts []inspector.Option
	if f.schema {
		opts = append(opts, inspector.WithSchemaAuto())
	}
	mod, err := inspector.Inspect(dir, opts...)
	if err != nil {
		return err
	}

	doc := render.Markdown(mod, file, cfg, displayVersion())

	if f.check {
		return check(file, mode, doc)
	}
	if err := output.Write(os.Stdout, file, mode, doc); err != nil {
		return err
	}
	if mode != config.ModeStdout && file != "" {
		fmt.Fprintf(os.Stderr, "wrote %s (%s)\n", file, mode)
	}
	return nil
}

// setRender turns on or off each section in a comma-separated list.
func setRender(cfg *config.Config, list string, render bool) error {
	for section := range strings.SplitSeq(list, ",") {
		if section = strings.TrimSpace(section); section == "" {
			continue
		}
		if err := cfg.SetRender(section, render); err != nil {
			return err
		}
	}
	return nil
}

// check returns an error when writing doc would change file.
func check(file, mode, doc string) error {
	want, err := output.Contents(file, mode, doc)
	if err != nil {
		return err
	}
	got, err := os.ReadFile(file)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if string(got) != want {
		return fmt.Errorf("%s is out of date; run open-docs without -check to update it", file)
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

// resolveOutput merges config output settings with the -m and -o flags. The
// flags, when set, win: -o sets the file (relative to the working directory)
// and forces replace mode unless -m picks another mode. Config file paths
// resolve relative to the module directory.
func resolveOutput(cfg *config.Config, dir, modeFlag, outFlag string) (file, mode string) {
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
	if modeFlag != "" {
		mode = modeFlag
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
