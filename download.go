package solc

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/lmittmann/go-solc/internal/mod"
	"golang.org/x/sync/singleflight"
)

var (
	maxTries = 2

	dg singleflight.Group // global download group
)

// checkSolc checks for the existence of the solc binary at
// "{MOD_ROOT}/.solc/bin/solc_{version}" and attempts to download it if it does
// not exist yet.
//
// The version string must be in the format "0.8.17".
func checkSolc(version SolcVersion) (string, error) {
	v, ok := solcVersions[version]
	if !ok {
		return "", fmt.Errorf("solc: unknown version %q", version)
	}

	if err := makeBinDir(); err != nil {
		return "", err
	}

	absSolcPath := filepath.Join(mod.Root, binPath, fmt.Sprintf("solc_v%s", version))

	_, err, _ := dg.Do(version.String(), func() (any, error) {
		if _, err := os.Stat(absSolcPath); errors.Is(err, os.ErrNotExist) {
			// download solc_{version}
			var (
				try int
				err error
			)
			for ; try < maxTries && (try == 0 || err != nil); try++ {
				err = downloadSolc(absSolcPath, v)
			}
			if try >= maxTries {
				return "", fmt.Errorf("solc: failed to download solc %q: %w", version, err)
			}
		}

		// solc_{version} binary exists
		f, err := os.Open(absSolcPath)
		if err != nil {
			return "", err
		}
		defer f.Close()

		return nil, verifyChecksum(version, f, v)
	})

	if err != nil {
		return "", err
	}
	return absSolcPath, nil
}

func verifyChecksum(version SolcVersion, r io.Reader, v solcVersion) error {
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
func downloadSolc(path string, v solcVersion) error {
	// request compiler
	resp, err := http.Get(solcBaseURL + v.Path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o0764)
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

// makeBinDir creates the directory ".solc/bin/" if it doesn't exist yet.
func makeBinDir() error {
	binDirMux.Lock()
	defer binDirMux.Unlock()

	return os.MkdirAll(filepath.Join(mod.Root, binPath), perm)
}

var binDirMux sync.Mutex
