package solc

import "testing"

func TestDownloadSolc(t *testing.T) {
	go func() {
		_, err := checkSolc("0.8.17")
		if err != nil {
			t.Log(err)
		}
	}()

	_, err := checkSolc("0.8.17")
	if err != nil {
		t.Fatal(err)
	}
}
