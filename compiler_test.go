package solc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCompile(t *testing.T) {
	c := &Compiler{Version: "0.8.17"}

	tests := []struct {
		Name         string
		ContractName string
		WantContract *Contract
		WantErr      error
	}{
		{
			Name:         "test1",
			ContractName: "Test1",
			WantContract: &Contract{
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
