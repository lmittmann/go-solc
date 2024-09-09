//go:generate go run params_gen.go
package solc

import (
	"github.com/lmittmann/go-solc/internal/version"
)

var (
	solcBaseURL  string
	solcVersions map[SolcVersion]solcVersion
)

// SolcVersion represents a solc version.
type SolcVersion string

func (v SolcVersion) String() string { return string(v) }

// Compare returns -1, 0, or +1 depending on whether v < other, v == other, or
// v > other
func (v SolcVersion) Cmp(other SolcVersion) int {
	return version.Compare(string(v), string(other))
}

// SolcVersions is a list of all available solc versions.
var SolcVersions []SolcVersion
