package solc_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lmittmann/solc"
)

func TestCompile(t *testing.T) {
	c := solc.New("0.8.17")

	tests := []struct {
		Name         string
		ContractName string
		WantContract *solc.Contract
		WantErr      error
	}{
		{
			Name:         "test1",
			ContractName: "Test1",
			WantContract: &solc.Contract{
				Code:       b("0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063f8a8fd6d14602d575b600080fd5b602a60405190815260200160405180910390f3fea26469706673582212209fce876a231310a49ee9848a2e6b937a9c3e07fb578272f974a37d4289160fee64736f6c63430008110033"),
				DeployCode: b("0x6080604052348015600f57600080fd5b50607780601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063f8a8fd6d14602d575b600080fd5b602a60405190815260200160405180910390f3fea26469706673582212209fce876a231310a49ee9848a2e6b937a9c3e07fb578272f974a37d4289160fee64736f6c63430008110033"),
			},
		},
		{
			Name:    "test2",
			WantErr: cmpopts.AnyError,
		},
	}

	for _, test := range tests {
		t.Run(test.Name+"_"+test.ContractName, func(t *testing.T) {
			contract, err := c.Compile("testdata/"+test.Name, test.ContractName)

			if diff := cmp.Diff(test.WantContract, contract); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
			if diff := cmp.Diff(test.WantErr, err,
				cmpopts.EquateErrors(),
			); diff != "" {
				t.Errorf("(-want, +got):\n%s", diff)
			}
		})
	}
}

// b returns a byte slice from a hexstring or panics if the hexstring does not
// represent a valid byte slice.
func b(hexBytes string) []byte {
	if !has0xPrefix(hexBytes) {
		panic(fmt.Sprintf("hex bytes %q must have 0x prefix", hexBytes))
	}
	if len(hexBytes)%2 != 0 {
		panic(fmt.Sprintf("hex bytes %q must have even number of hex chars", hexBytes))
	}

	bytes, err := hex.DecodeString(hexBytes[2:])
	if err != nil {
		panic(err)
	}
	return bytes
}

// has0xPrefix validates hexStr begins with '0x' or '0X'.
func has0xPrefix(hexStr string) bool {
	return len(hexStr) >= 2 && hexStr[0] == '0' && hexStr[1] == 'x'
}
