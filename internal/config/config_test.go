package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOverridesDefaults(t *testing.T) {
	path := filepath.Join(t.TempDir(), DefaultFilename)
	if err := os.WriteFile(path, []byte("main {\n  mode = \"inject\"\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	// mode comes from the user's file; docsfile falls back to defaults.hcl.
	want := MainConfig{DocsFile: "README.md", Mode: ModeInject}
	if cfg.Main == nil || *cfg.Main != want {
		t.Fatalf("Main = %+v, want %+v", cfg.Main, want)
	}
}
