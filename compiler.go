package solc

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lmittmann/go-solc/internal/console"
	"github.com/lmittmann/go-solc/internal/mod"
	"golang.org/x/sync/singleflight"
)

var (
	// The path within the module root where solc binaries are stored.
	binPath = ".solc/bin/"

	perm = os.FileMode(0o0775)

	// global compiler cache
	group    = new(singleflight.Group)
	cacheMux sync.RWMutex
	cache    = make(map[string]cacheItem)
)

type cacheItem struct {
	out *output
	err error
}

type Compiler struct {
	version SolcVersion // Solc version

	once        sync.Once
	solcAbsPath string // solc absolute path
	err         error  // initialization error
}

func New(version SolcVersion) *Compiler {
	return &Compiler{
		version: version,
	}
}

// init initializes the compiler.
func (c *Compiler) init() {
	// check mod root is set
	if mod.Root == "" {
		c.err = fmt.Errorf("solc: no go.mod detected")
		return
	}

	// check or download solc version
	c.solcAbsPath, c.err = checkSolc(c.version)
}

// Compile all contracts in the given directory and return the contract code of
// the contract with the given name.
func (c *Compiler) Compile(dir, contract string, opts ...Option) (*Contract, error) {
	out, err := c.compile(dir, contract, opts)
	if err != nil {
		return nil, err
	}

	// check for compilation errors
	if err := out.Err(); err != nil {
		return nil, err
	}

	// find contract code
	var con *Contract
	for _, conMap := range out.Contracts {
		for conName, c := range conMap {
			if conName == contract {
				con = &Contract{
					Runtime:     c.EVM.DeployedBytecode.Object,
					Constructor: c.EVM.Bytecode.Object,
					Code:        c.EVM.DeployedBytecode.Object,
					DeployCode:  c.EVM.Bytecode.Object,
				}
				break
			}
		}
	}
	if con == nil {
		return nil, fmt.Errorf("solc: unknown contract %q", contract)
	}
	return con, nil
}

// MustCompile is like [Compiler.Compile] but panics on error.
func (c *Compiler) MustCompile(dir, contract string, opts ...Option) *Contract {
	code, err := c.Compile(dir, contract, opts...)
	if err != nil {
		panic(err)
	}
	return code
}

// compile
func (c *Compiler) compile(baseDir, contract string, opts []Option) (*output, error) {
	// init an return on error
	c.once.Do(c.init)
	if c.err != nil {
		return nil, c.err
	}

	// check the directory exists
	if stat, err := os.Stat(baseDir); err != nil || !stat.IsDir() {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s is not a directory", baseDir)
	}

	// get absolute path of base directory
	absDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, err
	}

	// build src map
	srcMap, err := buildSrcMap(absDir)
	if err != nil {
		return nil, err
	}

	// add console.sol to src map
	srcMap["console.sol"] = src{
		Content: console.Src,
	}

	// build settings
	s := c.buildSettings(opts)
	in := &input{
		Lang:     s.lang,
		Sources:  srcMap,
		Settings: s,
	}

	// run solc
	return c.runWithCache(absDir, in)
}

func (c *Compiler) runWithCache(baseDir string, in *input) (*output, error) {
	// hash input
	h := sha256.New()
	if err := json.NewEncoder(h).Encode(in); err != nil {
		return nil, err
	}
	var hash [32]byte
	h.Sum(hash[:0])

	// run with cache
	cacheKey := fmt.Sprintf("%s_%x", c.version, hash)
	out, err, _ := group.Do(cacheKey, func() (any, error) {
		// check cache
		cacheMux.RLock()
		val, ok := cache[cacheKey]
		cacheMux.RUnlock()
		if ok {
			return val.out, val.err
		}

		// run solc
		out, err := c.run(baseDir, in)

		// update cache
		cacheMux.Lock()
		cache[cacheKey] = cacheItem{out, err}
		cacheMux.Unlock()

		return out, err
	})
	if err != nil {
		return nil, err
	}
	return out.(*output), nil
}

func (c *Compiler) run(baseDir string, in *input) (*output, error) {
	inputBuf := bytes.NewBuffer(nil)
	outputBuf := bytes.NewBuffer(nil)

	// encode input
	if err := json.NewEncoder(inputBuf).Encode(in); err != nil {
		return nil, err
	}

	var allowPaths []string
	allowPaths = append(allowPaths, baseDir)

	for _, remap := range in.Settings.Remappings {
		parts := strings.Split(remap, "=")
		if len(parts) != 2 {
			//invalid remapping
			continue
		}
		allowPaths = append(allowPaths, parts[1])
	}

	// run solc
	ex := exec.Command(c.solcAbsPath,
		"--allow-paths", strings.Join(allowPaths, ","),
		"--standard-json",
	)
	ex.Stdin = inputBuf
	ex.Stdout = outputBuf
	if err := ex.Run(); err != nil {
		return nil, err
	}

	// decode output
	var output *output
	if err := json.NewDecoder(outputBuf).Decode(&output); err != nil {
		return nil, err
	}
	return output, nil
}

func buildSrcMap(absDir string) (map[string]src, error) {
	fsys := os.DirFS(absDir)

	srcMap := make(map[string]src)
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(p) != ".sol" {
			return nil
		}
		srcMap[p] = src{
			URLS: []string{filepath.Join(absDir, p)},
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return srcMap, nil
}

// buildSettings builds the default settings and applies all options.
func (c *Compiler) buildSettings(opts []Option) *Settings {
	defaultEVMVersion, ok := defaultEVMVersions[c.version]
	if !ok {
		panic("unexpected solc version")
	}
	s := &Settings{
		lang:       defaultLang,
		Remappings: defaultRemappings,
		Optimizer:  defaultOptimizer,
		ViaIR:      defaultViaIR,
		EVMVersion: defaultEVMVersion,
	}
	for _, opt := range opts {
		opt(s)
	}
	s.OutputSelection = defaultOutputSelection
	return s
}
