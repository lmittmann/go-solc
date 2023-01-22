package internal

import (
	"strings"
	"testing"
)

func TestModRoot(t *testing.T) {
	if path := ModRoot(); !strings.HasSuffix(path, "solc/") {
		t.Fatalf("Unexpected workspace path: %q", path)
	}
}
