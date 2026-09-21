// Package output writes rendered documentation to its destination: stdout, a
// full file replace, or injection between markers in an existing file.
package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/remoterabbit/open-doc/internal/config"
)

// Markers delimit the open-doc managed region in inject mode. Everything
// between them (exclusive) is replaced on each run; the markers themselves are
// preserved so the region stays stable across runs.
const (
	BeginMarker = "<!-- BEGIN_OPEN_DOC -->"
	EndMarker   = "<!-- END_OPEN_DOC -->"
)

// Write dispatches on mode. In stdout mode (or when file is empty) it prints
// to w. Otherwise it writes to file using replace or inject semantics.
func Write(w *os.File, file, mode, content string) error {
	if mode == config.ModeStdout || file == "" {
		_, err := fmt.Fprint(w, content)
		return err
	}
	switch mode {
	case config.ModeReplace:
		return replace(file, content)
	case config.ModeInject:
		return inject(file, content)
	default:
		return fmt.Errorf("unknown output mode %q", mode)
	}
}

func replace(file, content string) error {
	return os.WriteFile(file, []byte(content), 0o644)
}

// inject replaces the text between BeginMarker and EndMarker in file with
// content, leaving the rest of the file untouched. If the file does not exist
// it is created with a single managed region. If the file exists but has no
// markers, a managed region is appended to the end.
func inject(file, content string) error {
	body := strings.TrimRight(content, "\n")

	existing, err := os.ReadFile(file)
	if errors.Is(err, fs.ErrNotExist) {
		return os.WriteFile(file, []byte(region(body)+"\n"), 0o644)
	}
	if err != nil {
		return err
	}

	updated, err := injectInto(string(existing), body)
	if err != nil {
		return err
	}
	return os.WriteFile(file, []byte(updated), 0o644)
}

// injectInto is the pure string transform behind inject, split out so it can
// be tested without touching the filesystem.
func injectInto(existing, body string) (string, error) {
	begin := strings.Index(existing, BeginMarker)
	end := strings.Index(existing, EndMarker)

	if begin == -1 && end == -1 {
		// No managed region yet: append one.
		out := existing
		if out != "" && !strings.HasSuffix(out, "\n") {
			out += "\n"
		}
		if out != "" {
			out += "\n"
		}
		return out + region(body) + "\n", nil
	}
	if begin == -1 || end == -1 || end < begin {
		return "", fmt.Errorf("malformed open-doc markers: need %q before %q", BeginMarker, EndMarker)
	}

	return existing[:begin] + region(body) + existing[end+len(EndMarker):], nil
}

// region wraps body in the begin/end markers.
func region(body string) string {
	return BeginMarker + "\n" + body + "\n" + EndMarker
}
