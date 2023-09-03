package mod_test

import (
	"strings"
	"testing"

	"github.com/lmittmann/go-solc/internal/mod"
)

func TestModRoot(t *testing.T) {
	if !strings.HasSuffix(mod.Root, "solc") {
		t.Fatalf("Unexpected module root: %q", mod.Root)
	}
}
