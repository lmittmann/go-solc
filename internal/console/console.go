//go:generate go run gen.go
package console

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Source of the console contracts.
//
//go:embed console.sol
var Src string

// Address of the console contract.
var Addr = common.HexToAddress("0x000000000000000000000000000000000baDC0DE")

var (
	argString  = abi.Argument{Type: abi.Type{T: abi.StringTy}}
	argUint    = abi.Argument{Type: abi.Type{T: abi.UintTy, Size: 256}}
	argInt     = abi.Argument{Type: abi.Type{T: abi.IntTy, Size: 256}}
	argBool    = abi.Argument{Type: abi.Type{T: abi.BoolTy}}
	argAddress = abi.Argument{Type: abi.Type{T: abi.AddressTy, Size: 20}}
	argBytes32 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}
	argBytes   = abi.Argument{Type: abi.Type{T: abi.BytesTy}}
)
