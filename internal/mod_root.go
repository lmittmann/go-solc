package internal

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

var (
	modRoot string
	once    sync.Once
)

// ModRoot returns the path to the root of the module.
func ModRoot() (path string) {
	once.Do(func() {
		stdout, _ := exec.Command("go", "env", "GOMOD").Output()
		modRoot = string(bytes.TrimSpace(stdout))
		if strings.HasSuffix(modRoot, "go.mod") {
			modRoot = strings.TrimSuffix(modRoot, "go.mod")
		} else {
			modRoot = ""
		}
	})
	return modRoot
}
