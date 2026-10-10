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
	Main *MainConfig `hcl:"main,block"`
}

type MainConfig struct {
	ConfigPath string `hcl:"configpath,optional"`
	DocsFile   string `hcl:"docsfile,optional"`
	Mode       string `hcl:"mode,optional"`
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

// Find returns the path to a .open-docs.hcl file, searching the module
// directory first and then the current working directory. The second return
// value reports whether a file was found.
func Find(moduleDir string) (string, bool) {
	candidates := []string{
		filepath.Join(moduleDir, DefaultFilename),
		DefaultFilename, // current working directory
	}
	for _, path := range candidates {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, true
		}
	}
	return "", false
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
		return fmt.Errorf("output.mode %q is invalid (want %q, %q, or %q)",
			c.Main.Mode, ModeStdout, ModeReplace, ModeInject)
	}
}
