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
var Addr = common.HexToAddress("0x000000000000000000636F6e736F6c652e6c6f67")

var (
	argString  = abi.Argument{Type: abi.Type{T: abi.StringTy}}
	argUint    = abi.Argument{Type: abi.Type{T: abi.UintTy, Size: 256}}
	argInt     = abi.Argument{Type: abi.Type{T: abi.IntTy, Size: 256}}
	argAddress = abi.Argument{Type: abi.Type{T: abi.AddressTy, Size: 20}}
	argBool    = abi.Argument{Type: abi.Type{T: abi.BoolTy}}
	argBytes   = abi.Argument{Type: abi.Type{T: abi.BytesTy}}
	argBytes1  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 1}}
	argBytes2  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 2}}
	argBytes3  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 3}}
	argBytes4  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 4}}
	argBytes5  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 5}}
	argBytes6  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 6}}
	argBytes7  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 7}}
	argBytes8  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 8}}
	argBytes9  = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 9}}
	argBytes10 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 10}}
	argBytes11 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 11}}
	argBytes12 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 12}}
	argBytes13 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 13}}
	argBytes14 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 14}}
	argBytes15 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 15}}
	argBytes16 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 16}}
	argBytes17 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 17}}
	argBytes18 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 18}}
	argBytes19 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 19}}
	argBytes20 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 20}}
	argBytes21 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 21}}
	argBytes22 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 22}}
	argBytes23 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 23}}
	argBytes24 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 24}}
	argBytes25 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 25}}
	argBytes26 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 26}}
	argBytes27 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 27}}
	argBytes28 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 28}}
	argBytes29 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 29}}
	argBytes30 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 30}}
	argBytes31 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 31}}
	argBytes32 = abi.Argument{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}
)
