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
	"os"
	"path/filepath"

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
	Main      *MainConfig      `hcl:"main,block"`
	Required  *RequiredConfig  `hcl:"required,block"`
	Providers *ProvidersConfig `hcl:"providers,block"`
	Variables *VariablesConfig `hcl:"variables,block"`
}

type MainConfig struct {
	DocsFile string `hcl:"docsfile,optional"` // file to write docs to
	Mode     string `hcl:"mode,optional"`     // one of three mode options: stdout, inject, or replace
	Header   bool   `hcl:"header,optional"`   // render the module title
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
