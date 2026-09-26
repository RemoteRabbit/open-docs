package render

import (
	"path/filepath"
	"testing"

	"github.com/remoterabbit/open-inspector/pkg/model"
)

func TestPositionLinksRelativeToOutputDocument(t *testing.T) {
	repo := t.TempDir()
	moduleDir := filepath.Join(repo, "examples", "vpc")
	p := model.Position{
		Filename: filepath.ToSlash(filepath.Join(moduleDir, "main file.tf")),
		Start:    model.Pos{Line: 4, Column: 5, Byte: 69},
		End:      model.Pos{Line: 7, Column: 6, Byte: 139},
	}

	got := position(p, moduleDir)
	want := "[`./main file.tf`](./main%20file.tf#L4-L7) [4:5:69 -> 7:6:139]"
	if got != want {
		t.Fatalf("position() = %q, want %q", got, want)
	}

	got = position(p, repo)
	want = "[`./examples/vpc/main file.tf`](./examples/vpc/main%20file.tf#L4-L7) [4:5:69 -> 7:6:139]"
	if got != want {
		t.Fatalf("position() = %q, want %q", got, want)
	}

	p.End.Line = p.Start.Line
	got = position(p, moduleDir)
	want = "[`./main file.tf`](./main%20file.tf#L4) [4:5:69 -> 4:6:139]"
	if got != want {
		t.Fatalf("position() = %q, want %q", got, want)
	}
}
