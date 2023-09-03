//go:generate go run gen.go
package console

import (
	_ "embed"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
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

// NewTracer returns a [vm.EVMLogger] that prints debug calls.
func NewTracer(tb testing.TB) vm.EVMLogger {
	tb.Helper()
	return &tracer{tb: tb}
}

type tracer struct {
	tb testing.TB
}

func (t *tracer) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	if to != Addr || len(input) < 4 {
		return
	}

	sel := ([4]byte)(input[:4])
	args, ok := Args[sel]
	if ok {
		t.logConsole(args, input[4:])
	}
}

func (*tracer) CaptureTxStart(uint64) {}
func (*tracer) CaptureTxEnd(uint64)   {}
func (*tracer) CaptureStart(*vm.EVM, common.Address, common.Address, bool, []byte, uint64, *big.Int) {
}
func (*tracer) CaptureEnd([]byte, uint64, error)  {}
func (*tracer) CaptureExit([]byte, uint64, error) {}
func (*tracer) CaptureState(uint64, vm.OpCode, uint64, uint64, *vm.ScopeContext, []byte, int, error) {
}
func (*tracer) CaptureFault(uint64, vm.OpCode, uint64, uint64, *vm.ScopeContext, int, error) {}

func (t *tracer) logConsole(args abi.Arguments, data []byte) {
	params, err := args.Unpack(data)
	if err != nil {
		t.tb.Fatalf("malformed log: %v", err)
	}

	for i, p := range params {
		switch p := p.(type) {
		case []byte:
			params[i] = hexutil.Bytes(p)
		}
	}

	t.tb.Log(params...)
}
