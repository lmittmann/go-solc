package solc

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/lmittmann/solc/debug"
	"github.com/lmittmann/solc/internal"
	"golang.org/x/sync/singleflight"
)

var (
	modRoot = internal.ModRoot()
	binPath = ".solc/bin/"

	perm = os.FileMode(0775)

	// global compilation cache
	group    = new(singleflight.Group)
	cacheMux sync.RWMutex
	cache    = make(map[string]cachItem)
)

type cachItem struct {
	out *output
	err error
}

type Compiler struct {
	Version string // Solc version

	once        sync.Once
	solcAbsPath string // solc absolute path
	err         error  // initialization error
}

// init initializes the compiler.
func (c *Compiler) init() {
	// check mod root is set
	if modRoot == "" {
		c.err = fmt.Errorf("solc: no go.mod detected")
		return
	}

	// create ".solc/bin/" dir if it doesn't exist
	if err := os.MkdirAll(modRoot+binPath, perm); err != nil {
		c.err = fmt.Errorf("solc: %w", err)
		return
	}

	// check or download solc version
	c.solcAbsPath, c.err = checkSolc(c.Version)
}

// Compile all contracts in the given directory and returns the contract code of
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
					Code:       c.EVM.DeployedBytecode.Object,
					DeployCode: c.EVM.Bytecode.Object,
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
		panic(fmt.Sprintf("solc: %v", err))
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
		return nil, fmt.Errorf("solc: %w", err)
	}

	// get absolute path of base directory
	absDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, err
	}

	// cache compilation run
	out, err, _ := group.Do(absDir, func() (any, error) {
		// check cache
		cacheMux.RLock()
		val, ok := cache[absDir]
		cacheMux.RUnlock()
		if ok {
			return val.out, val.err
		}

		// build src map
		srcMap, err := buildSrcMap(absDir)
		if err != nil {
			return nil, err
		}

		// add debug.sol to src map
		srcMap["debug.sol"] = src{
			Content: debug.Src,
		}

		// build settings
		s := buildSettings(opts)
		in := &input{
			Lang:     s.Lang,
			Sources:  srcMap,
			Settings: s,
		}

		// run solc
		out, err := c.run(absDir, in)

		// update cache
		cacheMux.Lock()
		cache[absDir] = cachItem{out, err}
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

	// run solc
	ex := exec.Command(c.solcAbsPath,
		"--allow-paths", strconv.Quote(baseDir),
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
func buildSettings(opts []Option) *Settings {
	s := &Settings{
		Lang:       defaultLang,
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
