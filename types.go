package solc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// Contract represents a compiled contract.
type Contract struct {
	Code       []byte // The bytecode of the contract after deployment.
	DeployCode []byte // The bytecode to deploy the contract.
}

// lang represents the language of the source code.
type lang string

const (
	langSolidity lang = "Solidity"
	langYul      lang = "Yul"
)

// EVMVersion represents the EVM version to compile for.
type EVMVersion string

const (
	EVMVersionPrague     EVMVersion = "prague"
	EVMVersionCancun     EVMVersion = "cancun"
	EVMVersionShanghai   EVMVersion = "shanghai"
	EVMVersionParis      EVMVersion = "paris"
	EVMVersionLondon     EVMVersion = "london"
	EVMVersionBerlin     EVMVersion = "berlin"
	EVMVersionIstanbul   EVMVersion = "istanbul"
	EVMVersionPetersburg EVMVersion = "petersburg"
	EVMVersionByzantium  EVMVersion = "byzantium"
)

type input struct {
	Lang     lang           `json:"language"`
	Sources  map[string]src `json:"sources"`
	Settings *Settings      `json:"settings"`
}

type src struct {
	Keccak256 string   `json:"keccak256,omitempty"`
	Content   string   `json:"content,omitempty"`
	URLS      []string `json:"urls,omitempty"`
}

// Settings for the compilation.
type Settings struct {
	lang            lang                           `json:"-"`
	Remappings      []string                       `json:"remappings,omitempty"`
	Optimizer       *Optimizer                     `json:"optimizer"`
	ViaIR           bool                           `json:"viaIR,omitempty"`
	EVMVersion      EVMVersion                     `json:"evmVersion"`
	OutputSelection map[string]map[string][]string `json:"outputSelection"`
}

type Optimizer struct {
	Enabled bool   `json:"enabled"`
	Runs    uint64 `json:"runs"`
}

type output struct {
	Errors    []error_                       `json:"errors"`
	Sources   map[string]srcOut              `json:"sources"`
	Contracts map[string]map[string]contract `json:"contracts"`
}

func (o *output) Err() error {
	var fmtMsgs []string
	for _, err := range o.Errors {
		if strings.EqualFold(err.Severity, "error") {
			fmtMsgs = append(fmtMsgs, err.FormattedMessage)
		}
	}

	if len(fmtMsgs) == 0 {
		return nil
	}
	return fmt.Errorf("solc: compilation failed\n%s", strings.Join(fmtMsgs, "\n"))
}

type error_ struct {
	SourceLocation   sourceLocation `json:"sourceLocation"`
	Type             string         `json:"type"`
	Component        string         `json:"component"`
	Severity         string         `json:"severity"`
	Message          string         `json:"message"`
	FormattedMessage string         `json:"formattedMessage"`
}

type sourceLocation struct {
	File  string `json:"file"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type srcOut struct {
	ID        int             `json:"id"`
	AST       json.RawMessage `json:"ast"`
	LegacyAST json.RawMessage `json:"legacyAST"`
}

type contract struct {
	ABI      []json.RawMessage `json:"abi"`
	Metadata string            `json:"metadata"`
	UserDoc  json.RawMessage   `json:"userdoc"`
	DevDoc   json.RawMessage   `json:"devdoc"`
	IR       string            `json:"ir"`
	EVM      evm               `json:"evm"`
}

type evm struct {
	Assembly          string                       `json:"assembly"`
	LegacyAssembly    json.RawMessage              `json:"legacyAssembly"`
	Bytecode          bytecode                     `json:"bytecode"`
	DeployedBytecode  bytecode                     `json:"deployedBytecode"`
	MethodIdentifiers map[string]string            `json:"methodIdentifiers"`
	GasEstimates      map[string]map[string]string `json:"gasEstimates"`
}

type bytecode struct {
	Object         hexBytes                              `json:"object"`
	Opcodes        string                                `json:"opcodes"`
	SourceMap      string                                `json:"sourceMap"`
	LinkReferences map[string]map[string][]linkReference `json:"linkReferences"`
}

type linkReference struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// hexBytes is a byte slice that is unmarshalled from a hexstring.
type hexBytes []byte

func (b *hexBytes) UnmarshalText(text []byte) error {
	*b = make([]byte, hex.DecodedLen(len(text)))
	_, err := hex.Decode(*b, text)
	return err
}

type solcVersion struct {
	Path   string
	Sha256 [32]byte
}
