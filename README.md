# `go-solc`: Go Bindings for the Solidity Compiler

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/go-solc.svg)](https://pkg.go.dev/github.com/lmittmann/go-solc)
[![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/go-solc)](https://goreportcard.com/report/github.com/lmittmann/go-solc)

`go-solc` provides an easy way to compile Solidity contracts from Go.

```
go get github.com/lmittmann/go-solc
```


## Getting Started

`go-solc` automatically downloads the specified version of the Solidity compiler from https://binaries.soliditylang.org and caches it at `.solc/bin/`.

Example test:
```go
// contract_test.go
func TestContract(t *testing.T) {
    c := solc.New(solc.VersionLatest)
    contract, err := c.Compile("src", "Test",
        solc.WithOptimizer(&solc.Optimizer{Enabled: true, Runs: 999999}),
    )
    // ...
}
```

Example directory structure:
```bash
workspace/
├── .solc/
│   └── bin/ # cached solc binaries
│       └── solc_v0.8.21
├── src/
│   └── test.sol
├── contract_test.go
├── go.mod
└── go.sum
```

> [!WARNING]
>
> This package is pre-1.0. There might be breaking changes between minor versions.
