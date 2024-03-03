package solc

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/lmittmann/go-solc/internal/console"
)

// NewConsole returns a [vm.EVMLogger] that logs calls of console.sol to the
// given testing.TB.
//
// To use console logging in your Solidity contract, import "console.sol".
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
		case []byte,
			[1]byte, [2]byte, [3]byte, [4]byte, [5]byte, [6]byte, [7]byte, [8]byte,
			[9]byte, [10]byte, [11]byte, [12]byte, [13]byte, [14]byte, [15]byte, [16]byte,
			[17]byte, [18]byte, [19]byte, [20]byte, [21]byte, [22]byte, [23]byte, [24]byte,
			[25]byte, [26]byte, [27]byte, [28]byte, [29]byte, [30]byte, [31]byte, [32]byte:
			strVals[i] = fmt.Sprintf("0x%x", p)
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
