//go:build !(linux || darwin)

package solc

var solcBaseURL = ""

var solcVersions = map[string]solcVersion{}
