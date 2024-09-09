//go:build ignore

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/lmittmann/go-solc/internal/version"
)

var (
	solcBaseURL = "https://binaries.soliditylang.org/"

	tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"replaceAll": strings.ReplaceAll,
		"last":       func(s []*build) int { return len(s) - 1 },
	}).Parse(tmplStr))
)

func main() {
	targets := []*target{
		{
			BaseURL:    solcBaseURL + "linux-amd64/",
			Fn:         "params_linux_amd64.go",
			MinVersion: "0.5.0",
		},
		{
			BaseURL:    solcBaseURL + "macosx-amd64/",
			Fn:         "params_darwin_amd64.go",
			MinVersion: "0.5.0",
		},
		{
			BaseURL:    solcBaseURL + "macosx-amd64/",
			Fn:         "params_darwin_arm64.go",
			MinVersion: "0.8.24",
		},
	}

	errCh := make(chan error)
	for _, target := range targets {
		target := target

		go func() { errCh <- gen(target) }()
	}

	for i := 0; i < len(targets); i++ {
		if err := <-errCh; err != nil {
			fmt.Println(err)
		}
	}
}

func gen(target *target) error {
	// fetch and decode list.json
	resp, err := http.Get(target.BaseURL + "list.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var list struct {
		Builds []*build `json:"builds"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return err
	}

	// open file
	f, err := os.Create(target.Fn)
	if err != nil {
		return err
	}
	defer f.Close()

	// execute template
	model := &model{
		Target: target,
		Builds: slices.DeleteFunc(list.Builds, func(build *build) bool {
			return version.Compare(build.Version, target.MinVersion) < 0
		}),
	}
	if err := tmpl.Execute(f, model); err != nil {
		return err
	}
	return nil
}

func parseVersion(v string) (major, minor, patch int, err error) {
	_, err = fmt.Sscanf(v, "%d.%d.%d", &major, &minor, &patch)
	return
}

type target struct {
	BaseURL    string
	Fn         string
	MinVersion string
}

type build struct {
	Path    string  `json:"path"`
	Version string  `json:"version"`
	Sha256  bytes32 `json:"sha256"`
}

// bytes32 is a byte array that is unmarshalled from a hexstring.
type bytes32 [32]byte

func (b *bytes32) UnmarshalText(text []byte) error {
	if len(text) != 66 {
		return fmt.Errorf("invalid hex string")
	}
	_, err := hex.Decode(b[:], text[2:])
	return err
}

type model struct {
	Target *target  `json:"target"`
	Builds []*build `json:"builds"`
}

const tmplStr = `// Code generated by "go generate"; DO NOT EDIT.

package solc

func init() {
	solcBaseURL = "{{ .Target.BaseURL }}"

	solcVersions = map[SolcVersion]solcVersion{
	{{- range .Builds }}
		{{ $version := (printf "SolcVersion%s:" (replaceAll .Version "." "_")) -}}
		{{ printf "%-18s" $version }} {Sha256: [32]byte{
		{{- range $i, $elem := .Sha256 -}}
			{{- if $i }}, {{ end -}}
			{{- printf "%#02v" $elem -}}
		{{- end -}}
		}, Path: "{{ .Path }}"},
	{{- end }}
	}

	SolcVersions = []SolcVersion{
	{{- range .Builds }}
		{{ printf "SolcVersion%s," (replaceAll .Version "." "_") }}
	{{- end }}
	}
}

const (
{{- range .Builds }}
	{{ $version := (printf "%s" .Version) -}}
	{{ printf "SolcVersion%-6s SolcVersion = %q" (replaceAll .Version "." "_") .Version }}
{{- end }}

	// Latest version of solc.
	SolcVersionLatest = SolcVersion{{ replaceAll (index .Builds (last .Builds)).Version "." "_" }}
)
`
