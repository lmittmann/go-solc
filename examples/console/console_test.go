package console_test

import (
	"fmt"
	"testing"

	"github.com/lmittmann/solc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/lmittmann/w3/w3vm"
)

var (
	c = solc.Compiler{Version: "0.8.21"}
)

func TestNewConsole(t *testing.T) {
	// compile contract
	contract, err := c.Compile(".", "ConsoleTest")
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}
	contractAddr := w3vm.RandA()

	tests := []struct {
		Name     string
		WantLogs []string
	}{
		{
			Name:     "testHello",
			WantLogs: []string{"Hello, World!"},
		},
		{
			Name:     "testInt",
			WantLogs: []string{"-42"},
		},
		{
			Name:     "testBytes",
			WantLogs: []string{"0x74657374"},
		},
		{
			Name:     "testBytes1",
			WantLogs: []string{"0xff"},
		},
		{
			Name:     "testAddress",
			WantLogs: []string{"0x000000000000000000000000000000000000c0Fe"},
		},
		{
			Name:     "testBool",
			WantLogs: []string{"true"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// construct test function binding
			funcTest := w3.MustNewFunc(test.Name+"()", "")

			// setup test VM
			vm, _ := w3vm.New(
				w3vm.WithState(w3types.State{
					contractAddr: {
						Code: contract.Code,
					},
				}),
			)

			// apply test function
			if _, err := vm.Apply(
				&w3types.Message{
					To:   &contractAddr,
					Func: funcTest,
				},
				solc.NewConsole(newMockTB(t, test.WantLogs)),
			); err != nil {
				t.Fatal(err)
			}
		})
	}
}

type mockTB struct {
	testing.TB

	i        int
	wantLogs []string
}

func newMockTB(tb testing.TB, wantLogs []string) *mockTB {
	return &mockTB{
		TB:       tb,
		wantLogs: wantLogs,
	}
}

func (tb *mockTB) Log(args ...interface{}) {
	defer func() { tb.i++ }()

	got := fmt.Sprint(args...)

	if tb.i >= len(tb.wantLogs) {
		tb.Fatalf("unexpected log: %q", got)
	}
	if want := tb.wantLogs[tb.i]; want != got {
		tb.Fatalf("Log[%d]: want %q, got %q", tb.i, want, got)
	}
}
