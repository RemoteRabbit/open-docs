package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindWalksUpToRepoRoot(t *testing.T) {
	// outer/.open-docs.hcl must never be found: outer/repo is the repo root.
	outer := t.TempDir()
	repo := filepath.Join(outer, "repo")
	mod := filepath.Join(repo, "modules", "vpc")
	for _, dir := range []string{filepath.Join(repo, ".git"), mod} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	write := func(path string) {
		if err := os.WriteFile(path, nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write(filepath.Join(outer, DefaultFilename))

	if got, ok := Find(mod); ok {
		t.Fatalf("Find found %s above the repo root", got)
	}

	rootCfg := filepath.Join(repo, DefaultFilename)
	write(rootCfg)
	if got, ok := Find(mod); !ok || got != rootCfg {
		t.Fatalf("Find = %q, %v; want repo root config %q", got, ok, rootCfg)
	}

	modCfg := filepath.Join(mod, DefaultFilename)
	write(modCfg)
	if got, ok := Find(mod); !ok || got != modCfg {
		t.Fatalf("Find = %q, %v; want nearest config %q", got, ok, modCfg)
	}
}

func TestLoadOverridesDefaults(t *testing.T) {
	path := filepath.Join(t.TempDir(), DefaultFilename)
	if err := os.WriteFile(path, []byte("main {\n  mode = \"inject\"\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	// mode comes from the user's file; docsfile and header fall back to defaults.hcl.
	want := MainConfig{DocsFile: "README.md", Mode: ModeInject, Header: true}
	if cfg.Main == nil || *cfg.Main != want {
		t.Fatalf("Main = %+v, want %+v", cfg.Main, want)
	}
}
