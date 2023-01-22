# `solc`: Go Bindings for the Solidity Compiler

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/solc.svg)](https://pkg.go.dev/github.com/lmittmann/solc)
[![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/solc)](https://goreportcard.com/report/github.com/lmittmann/solc)

`solc` provides an easy way to compile Solidity contracts from Go.

```
go get github.com/lmittmann/solc
```


## Getting Started

`solc` will automatically download the specified version of the Solidity compiler from [`binaries.soliditylang.org`](https://binaries.soliditylang.org) and caches it at `.solc/bin/`.

Example test that will panic if the contract cannot be compiled:
```go
// contract_test.go
var (
    c        = &solc.Compiler{Version: "0.8.17"}
    contract = c.MustCompile("src", "Test")
)

func TestContract(t *testing.T) {
    // ...
}
```

Example directory structure:
```bash
workspace/
├── .solc/
│   └── bin/ # cached solc binaries
│       └── solc_v0.8.17
├── src/
│   └── test.sol
├── contract_test.go
├── go.mod
└── go.sum
```

> **Warning**
>
> This package is pre-1.0. There might be breaking changes between minor versions.
