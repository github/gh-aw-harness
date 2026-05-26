package harness

import "testing"

func TestName(t *testing.T) {
	if got := Name(); got != "gh-aw-harness" {
		t.Fatalf("Name() = %q, want %q", got, "gh-aw-harness")
	}
}
