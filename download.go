package solc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lmittmann/solc/internal/mod"
	"golang.org/x/sync/singleflight"
)

var (
	maxTries = 2

	dg singleflight.Group // global download group
)

// checkSolc checks for the existence of the solc binary at
// "{MOD_ROOT}/.solc/bin/solc_{version}" and atempts to download it if it does
// not exist yet.
//
// The version string must be in the format "0.8.17".
func checkSolc(version string) (string, error) {
	v, ok := solcVersions[version]
	if !ok {
		return "", fmt.Errorf("solc: unknown version %q", version)
	}

	absSolcPath := filepath.Join(mod.Root, binPath, fmt.Sprintf("solc_v%s", version))

	f, err := os.Open(absSolcPath)
	if err == nil {
		// solc_{version} binary exists
		defer f.Close()
		return absSolcPath, verifyChecksum(version, f, v)
	}

	// download solc_{version}
	for try := 0; try < maxTries; try++ {
		_, err, _ = dg.Do(version, func() (any, error) {
			return nil, downloadSolc(version, absSolcPath, v)
		})
		if err == nil {
			break
		}
	}
	if err != nil {
		return "", err
	}

	// solc_{version} binary exists
	f, err = os.Open(absSolcPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return absSolcPath, verifyChecksum(version, f, v)

}

func verifyChecksum(version string, r io.Reader, v solcVersion) error {
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return err
	}

	var gotSha256 [32]byte
	hash.Sum(gotSha256[:0])

	if v.Sha256 != gotSha256 {
		return fmt.Errorf("solc: checksum mismatch for version %q", version)
	}
	return nil
}

// downloadSolc downloads the solc binary with the given version, writes it to a
// file at the given path and return its content.
func downloadSolc(version, path string, v solcVersion) error {
	// request compiler
	resp, err := http.Get(solcBaseURL + v.Path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0764)
	if err != nil {
		return err
	}
	defer f.Close()

	// copy response body to file
	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}
	return nil
}
