package output

import "testing"

func TestInjectIntoReplacesBetweenMarkers(t *testing.T) {
	existing := "# Title\n\nintro\n\n" + BeginMarker + "\nold content\n" + EndMarker + "\n\nfooter\n"
	got, err := injectInto(existing, "new content")
	if err != nil {
		t.Fatalf("injectInto: %v", err)
	}
	want := "# Title\n\nintro\n\n" + BeginMarker + "\nnew content\n" + EndMarker + "\n\nfooter\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}

func TestInjectIntoAppendsWhenNoMarkers(t *testing.T) {
	got, err := injectInto("# Title\n\nintro\n", "body")
	if err != nil {
		t.Fatalf("injectInto: %v", err)
	}
	want := "# Title\n\nintro\n\n" + BeginMarker + "\nbody\n" + EndMarker + "\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}

func TestInjectIntoEmptyFile(t *testing.T) {
	got, err := injectInto("", "body")
	if err != nil {
		t.Fatalf("injectInto: %v", err)
	}
	want := BeginMarker + "\nbody\n" + EndMarker + "\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}

func TestInjectIntoMalformedMarkers(t *testing.T) {
	// End before begin.
	existing := EndMarker + "\n" + BeginMarker + "\n"
	if _, err := injectInto(existing, "body"); err == nil {
		t.Errorf("expected error for markers in wrong order")
	}
}
