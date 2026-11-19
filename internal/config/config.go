// Package config loads the optional .open-docs.hcl configuration file.
//
// Decoding is static (Level 1): the file is parsed with a nil evaluation
// context, so no variables or functions are available yet. The dynamic eval
// context (for/conditionals reacting to the inspected module) is a later
// phase.
package config

import (
	_ "embed"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// DefaultFilename is the config file name open-docs looks for.
const DefaultFilename = ".open-docs.hcl"

// Output modes.
const (
	ModeStdout  = "stdout"  // print to stdout
	ModeReplace = "replace" // overwrite the whole target file
	ModeInject  = "inject"  // replace content between markers in the target file
)

// Config is the decoded .open-docs.hcl file. All blocks are optional so an
// empty or absent file is valid; anything it leaves unset keeps the value
// from defaults.hcl.
type Config struct {
	Main        *MainConfig        `hcl:"main,block"`
	Required    *RequiredConfig    `hcl:"required,block"`
	Providers   *ProvidersConfig   `hcl:"providers,block"`
	Variables   *VariablesConfig   `hcl:"variables,block"`
	Outputs     *OutputsConfig     `hcl:"outputs,block"`
	Locals      *LocalsConfig      `hcl:"locals,block"`
	Resources   *ResourcesConfig   `hcl:"resources,block"`
	Data        *ResourcesConfig   `hcl:"data,block"`
	Ephemeral   *ResourcesConfig   `hcl:"ephemeral,block"`
	Modules     *ModulesConfig     `hcl:"modules,block"`
	Diagnostics *DiagnosticsConfig `hcl:"diagnostics,block"`
	Moved       *MovedConfig       `hcl:"moved,block"`
	Imports     *ImportsConfig     `hcl:"imports,block"`
	Removed     *RemovedConfig     `hcl:"removed,block"`
	Checks      *ChecksConfig      `hcl:"checks,block"`
	Schema      *SchemaConfig      `hcl:"schema_findings,block"`
}

type MainConfig struct {
	DocsFile string `hcl:"docsfile,optional"` // file to write docs to
	Mode     string `hcl:"mode,optional"`     // one of three mode options: stdout, inject, or replace
	Header   bool   `hcl:"header,optional"`   // render the module title
	Verbose  bool   `hcl:"verbose,optional"`  // render full byte positions
	Version  bool   `hcl:"version,optional"`
}

type RequiredConfig struct {
	Render   bool `hcl:"render,optional"`
	Name     bool `hcl:"name,optional"`
	Source   bool `hcl:"source,optional"`
	Version  bool `hcl:"version,optional"`
	Alias    bool `hcl:"alias,optional"`
	Position bool `hcl:"position,optional"`
}

type ProvidersConfig struct {
	Render   bool `hcl:"render,optional"`
	Name     bool `hcl:"name,optional"`
	Alias    bool `hcl:"alias,optional"`
	ForEach  bool `hcl:"for_each,optional"`
	Position bool `hcl:"position,optional"`
}

type VariablesConfig struct {
	Render      bool `hcl:"render,optional"`
	Name        bool `hcl:"name,optional"`
	Type        bool `hcl:"type,optional"`
	Default     bool `hcl:"default,optional"`
	Required    bool `hcl:"required,optional"`
	Description bool `hcl:"description,optional"`
	Comment     bool `hcl:"comment,optional"`
	Sensitive   bool `hcl:"sensitive,optional"`
	Nullable    bool `hcl:"nullable,optional"`
	Ephemeral   bool `hcl:"ephemeral,optional"`
	Validation  bool `hcl:"validation,optional"`
	Position    bool `hcl:"position,optional"`
}

type OutputsConfig struct {
	Render      bool `hcl:"render,optional"`
	Name        bool `hcl:"name,optional"`
	Value       bool `hcl:"value,optional"`
	References  bool `hcl:"references,optional"`
	Description bool `hcl:"description,optional"`
	Comment     bool `hcl:"comment,optional"`
	Sensitive   bool `hcl:"sensitive,optional"`
	Ephemeral   bool `hcl:"ephemeral,optional"`
	DependsOn   bool `hcl:"depends_on,optional"`
	Position    bool `hcl:"position,optional"`
}

type LocalsConfig struct {
	Render     bool `hcl:"render,optional"`
	Name       bool `hcl:"name,optional"`
	Value      bool `hcl:"value,optional"`
	References bool `hcl:"references,optional"`
	Position   bool `hcl:"position,optional"`
}

type ResourcesConfig struct {
	Render     bool `hcl:"render,optional"`
	Name       bool `hcl:"name,optional"`
	Mode       bool `hcl:"mode,optional"`
	Type       bool `hcl:"type,optional"`
	Provider   bool `hcl:"provider,optional"`
	Count      bool `hcl:"count,optional"`
	ForEach    bool `hcl:"for_each,optional"`
	DependsOn  bool `hcl:"depends_on,optional"`
	NestedBody bool `hcl:"nested_body,optional"`
	Lifecycle  bool `hcl:"lifecycle,optional"`
	Comment    bool `hcl:"comment,optional"`
	Position   bool `hcl:"position,optional"`
}

type ModulesConfig struct {
	Render    bool `hcl:"render,optional"`
	Name      bool `hcl:"name,optional"`
	Source    bool `hcl:"source,optional"`
	Version   bool `hcl:"version,optional"`
	Count     bool `hcl:"count,optional"`
	ForEach   bool `hcl:"for_each,optional"`
	DependsOn bool `hcl:"depends_on,optional"`
	Providers bool `hcl:"providers,optional"`
	Inputs    bool `hcl:"inputs,optional"`
	Comment   bool `hcl:"comment,optional"`
	Position  bool `hcl:"position,optional"`
}

type DiagnosticsConfig struct {
	Render   bool `hcl:"render,optional"`
	Severity bool `hcl:"severity,optional"`
	Summary  bool `hcl:"summary,optional"`
	Detail   bool `hcl:"detail,optional"`
	Subject  bool `hcl:"subject,optional"`
	Context  bool `hcl:"context,optional"`
}

type MovedConfig struct {
	Render   bool `hcl:"render,optional"`
	From     bool `hcl:"from,optional"`
	To       bool `hcl:"to,optional"`
	Position bool `hcl:"position,optional"`
}

type ImportsConfig struct {
	Render   bool `hcl:"render,optional"`
	To       bool `hcl:"to,optional"`
	Id       bool `hcl:"id,optional"`
	Provider bool `hcl:"provider,optional"`
	Position bool `hcl:"position,optional"`
}

type RemovedConfig struct {
	Render        bool `hcl:"render,optional"`
	From          bool `hcl:"from,optional"`
	DestroyOnDrop bool `hcl:"destroy_on_drop,optional"`
	Position      bool `hcl:"position,optional"`
}

type ChecksConfig struct {
	Render     bool `hcl:"render,optional"`
	Name       bool `hcl:"name,optional"`
	DataSource bool `hcl:"data_source,optional"`
	Assertions bool `hcl:"assertions,optional"`
	Position   bool `hcl:"position,optional"`
}

type SchemaConfig struct {
	Render    bool `hcl:"render,optional"`
	Resource  bool `hcl:"resource,optional"`
	Kind      bool `hcl:"kind,optional"`
	Attribute bool `hcl:"attribute,optional"`
	Message   bool `hcl:"message,optional"`
	Position  bool `hcl:"position,optional"`
}

//go:embed defaults.hcl
var defaultsHCL []byte

// Default returns a Config with the built-in defaults from defaults.hcl.
func Default() *Config {
	var cfg Config
	if err := hclsimple.Decode("defaults.hcl", defaultsHCL, nil, &cfg); err != nil {
		panic(fmt.Sprintf("config: invalid embedded defaults.hcl: %v", err))
	}
	return &cfg
}

// Find returns the path to the nearest .open-docs.hcl file, searching the
// module directory and then each parent up to the repository root (the first
// directory containing .git) or the filesystem root. The second return value
// reports whether a file was found.
func Find(moduleDir string) (string, bool) {
	dir, err := filepath.Abs(moduleDir)
	if err != nil {
		return "", false
	}
	for {
		path := filepath.Join(dir, DefaultFilename)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, true
		}
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return "", false
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

// Load reads and decodes the config file at path on top of Default(), so
// attributes the file omits keep their default values.
func Load(path string) (*Config, error) {
	cfg := Default()
	if err := hclsimple.DecodeFile(path, nil, cfg); err != nil {
		return nil, fmt.Errorf("config %s: %w", path, err)
	}
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config %s: %w", path, err)
	}
	return cfg, nil
}

// SetRender turns a section on or off by its block name, such as "variables"
// or "schema_findings", as the -show and -hide flags do.
func (c *Config) SetRender(section string, render bool) error {
	switches := map[string]*bool{
		"required":        &c.Required.Render,
		"providers":       &c.Providers.Render,
		"variables":       &c.Variables.Render,
		"outputs":         &c.Outputs.Render,
		"locals":          &c.Locals.Render,
		"resources":       &c.Resources.Render,
		"data":            &c.Data.Render,
		"ephemeral":       &c.Ephemeral.Render,
		"modules":         &c.Modules.Render,
		"diagnostics":     &c.Diagnostics.Render,
		"moved":           &c.Moved.Render,
		"imports":         &c.Imports.Render,
		"removed":         &c.Removed.Render,
		"checks":          &c.Checks.Render,
		"schema_findings": &c.Schema.Render,
	}
	sw, ok := switches[section]
	if !ok {
		names := slices.Sorted(maps.Keys(switches))
		return fmt.Errorf("unknown section %q (want one of %s)", section, strings.Join(names, ", "))
	}
	*sw = render
	return nil
}

func (c *Config) validate() error {
	if c.Main == nil {
		return nil
	}

	switch c.Main.Mode {
	case "", ModeStdout, ModeReplace, ModeInject:
		return nil
	default:
		return fmt.Errorf("main.mode %q is invalid (want %q, %q, or %q)",
			c.Main.Mode, ModeStdout, ModeReplace, ModeInject)
	}
}
