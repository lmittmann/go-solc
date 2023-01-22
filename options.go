package solc

// default settings options
var (
	defaultLang                     = LangSolidity
	defaultRemappings      []string = nil
	defaultOptimizer                = &Optimizer{Enabled: true, Runs: 200}
	defaultViaIR                    = false
	defaultEVMVersion               = EVMVersionLondon
	defaultOutputSelection          = map[string]map[string][]string{
		"*": {
			"*": {"evm.bytecode.object", "evm.deployedBytecode.object"},
		},
	}
)

// An Option configures the compilation [Settings].
type Option func(*Settings)

// WithOptimizer configures the compilation [Settings] to set the given
// [Optimizer].
func WithOptimizer(o *Optimizer) Option {
	return func(s *Settings) {
		s.Optimizer = o
	}
}

// WithViaIR configures the compilation [Settings] to set viaIR to the given
// parameter "enabled".
func WithViaIR(enabled bool) Option {
	return func(s *Settings) {
		s.ViaIR = enabled
	}
}

// WithEVMVersion configures the compilation [Settings] to set the given EVM
// version.
func WithEVMVersion(evmVersion EVMVersion) Option {
	return func(s *Settings) {
		s.EVMVersion = evmVersion
	}
}
