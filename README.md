# `solc`: Go Bindings for the Solidity Compiler

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/solc.svg)](https://pkg.go.dev/github.com/lmittmann/solc)
[![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/solc)](https://goreportcard.com/report/github.com/lmittmann/solc)

`solc` provides an easy way to compile Solidity contracts from Go.

```
go get github.com/lmittmann/solc
```


## Getting Started

`solc` automatically downloads the specified version of the Solidity compiler from [`binaries.soliditylang.org`](https://binaries.soliditylang.org) and caches it at `.solc/bin/`.

Example test:
```go
// contract_test.go
func TestContract(t *testing.T) {
    c := solc.New("0.8.21")
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

> **Warning**
>
> This package is pre-1.0. There might be breaking changes between minor versions.
