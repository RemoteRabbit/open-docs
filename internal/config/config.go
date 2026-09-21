// Package config loads the optional .open-doc.hcl configuration file.
//
// Decoding is static (Level 1): the file is parsed with a nil evaluation
// context, so no variables or functions are available yet. The dynamic eval
// context (for/conditionals reacting to the inspected module) is a later
// phase; see ROADMAP.md.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// DefaultFilename is the config file name open-doc looks for.
const DefaultFilename = ".open-doc.hcl"

// Output modes.
const (
	ModeStdout  = "stdout"  // print to stdout
	ModeReplace = "replace" // overwrite the whole target file
	ModeInject  = "inject"  // replace content between markers in the target file
)

// Config is the decoded .open-doc.hcl file. All blocks are optional so an
// empty or absent file is valid and yields zero values (defaults applied by
// the caller).
type Config struct {
	// TODO: Expand this to handle more blocks of differing types also would need to update main.tf
	Output *OutputConfig `hcl:"file,block"`
}

// OutputConfig controls where and how rendered docs are written.
type OutputConfig struct {
	// File is the target path. Relative paths resolve against the module
	// directory. Empty means stdout.
	File string `hcl:"file,optional"`
	// Mode is one of "stdout", "replace", or "inject". Empty defaults to
	// "replace" when File is set, otherwise "stdout".
	Mode string `hcl:"mode,optional"`
}

// Default returns a Config with the built-in defaults (stdout output).
func Default() *Config {
	return &Config{}
}

// Find returns the path to a .open-doc.hcl file, searching the module
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

// Load reads and decodes the config file at path.
func Load(path string) (*Config, error) {
	var cfg Config
	if err := hclsimple.DecodeFile(path, nil, &cfg); err != nil {
		return nil, fmt.Errorf("config %s: %w", path, err)
	}
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config %s: %w", path, err)
	}
	return &cfg, nil
}

func (c *Config) validate() error {
	if c.Output == nil {
		return nil
	}
	switch c.Output.Mode {
	case "", ModeStdout, ModeReplace, ModeInject:
		return nil
	default:
		return fmt.Errorf("output.mode %q is invalid (want %q, %q, or %q)",
			c.Output.Mode, ModeStdout, ModeReplace, ModeInject)
	}
}
