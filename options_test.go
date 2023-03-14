package solc

import "testing"

func TestDefaultEVMVersions(t *testing.T) {
	if len(solcVersions) == 0 {
		t.Fatalf(`empty "solcVersions". unsupported build target?`)
	}

	for version := range solcVersions {
		if _, ok := defaultEVMVersions[version]; !ok {
			t.Errorf("Missing default EVM version for %q", version)
		}
	}
}
