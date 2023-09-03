package solc

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/lmittmann/solc/internal/console"
)

// NewConsole returns a new [vm.EVMLogger] that logs calls to console.sol to the
// given testing.TB.
func NewConsole(tb testing.TB) vm.EVMLogger {
	return &consoleTracer{tb: tb}
}

type consoleTracer struct {
	tb testing.TB
}

func (ct *consoleTracer) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	if to != console.Addr || len(input) < 4 {
		return
	}

	sel := ([4]byte)(input[:4])
	args, ok := console.Args[sel]
	if ok {
		ct.log(args, input[4:])
	}
}

func (ct *consoleTracer) log(args abi.Arguments, data []byte) {
	params, err := args.Unpack(data)
	if err != nil {
		ct.tb.Fatalf("malformed log: %v", err)
		return
	}

	strVals := make([]string, len(params))
	for i, p := range params {
		switch p := p.(type) {
		case []byte:
			strVals[i] = fmt.Sprintf("%x", p)
		default:
			strVals[i] = fmt.Sprint(p)
		}
	}
	ct.tb.Log(strings.Join(strVals, " "))
}

func (*consoleTracer) CaptureTxStart(uint64) {}
func (*consoleTracer) CaptureTxEnd(uint64)   {}
func (*consoleTracer) CaptureStart(*vm.EVM, common.Address, common.Address, bool, []byte, uint64, *big.Int) {
}
func (*consoleTracer) CaptureEnd([]byte, uint64, error)  {}
func (*consoleTracer) CaptureExit([]byte, uint64, error) {}
func (*consoleTracer) CaptureState(uint64, vm.OpCode, uint64, uint64, *vm.ScopeContext, []byte, int, error) {
}
func (*consoleTracer) CaptureFault(uint64, vm.OpCode, uint64, uint64, *vm.ScopeContext, int, error) {
}
