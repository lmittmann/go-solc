package solc

import "testing"

func TestDownloadSolc(t *testing.T) {
	errCh := make(chan error, 2)
	go func() {
		_, err := checkSolc("0.8.21")
		errCh <- err
	}()

	_, err := checkSolc("0.8.21")
	errCh <- err

	for i := 0; i < cap(errCh); i++ {
		if err := <-errCh; err != nil {
			t.Fatal(err)
		}
	}
}
