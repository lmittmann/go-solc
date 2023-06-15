package solc

// default settings options
var (
	defaultLang                     = langSolidity
	defaultRemappings      []string = nil
	defaultOptimizer                = &Optimizer{Enabled: true, Runs: 200}
	defaultViaIR                    = false
	defaultOutputSelection          = map[string]map[string][]string{
		"*": {
			"*": {"evm.bytecode.object", "evm.deployedBytecode.object"},
		},
	}
	defaultEVMVersions = map[string]EVMVersion{
		"0.5.0":  EVMVersionByzantium,
		"0.5.1":  EVMVersionByzantium,
		"0.5.2":  EVMVersionByzantium,
		"0.5.3":  EVMVersionByzantium,
		"0.5.4":  EVMVersionByzantium,
		"0.5.5":  EVMVersionPetersburg,
		"0.5.6":  EVMVersionPetersburg,
		"0.5.7":  EVMVersionPetersburg,
		"0.5.8":  EVMVersionPetersburg,
		"0.5.9":  EVMVersionPetersburg,
		"0.5.10": EVMVersionPetersburg,
		"0.5.11": EVMVersionPetersburg,
		"0.5.12": EVMVersionPetersburg,
		"0.5.13": EVMVersionPetersburg,
		"0.5.14": EVMVersionIstanbul,
		"0.5.15": EVMVersionIstanbul,
		"0.5.16": EVMVersionIstanbul,
		"0.5.17": EVMVersionIstanbul,
		"0.6.0":  EVMVersionIstanbul,
		"0.6.1":  EVMVersionIstanbul,
		"0.6.2":  EVMVersionIstanbul,
		"0.6.3":  EVMVersionIstanbul,
		"0.6.4":  EVMVersionIstanbul,
		"0.6.5":  EVMVersionIstanbul,
		"0.6.6":  EVMVersionIstanbul,
		"0.6.7":  EVMVersionIstanbul,
		"0.6.8":  EVMVersionIstanbul,
		"0.6.9":  EVMVersionIstanbul,
		"0.6.10": EVMVersionIstanbul,
		"0.6.11": EVMVersionIstanbul,
		"0.6.12": EVMVersionIstanbul,
		"0.7.0":  EVMVersionIstanbul,
		"0.7.1":  EVMVersionIstanbul,
		"0.7.2":  EVMVersionIstanbul,
		"0.7.3":  EVMVersionIstanbul,
		"0.7.4":  EVMVersionIstanbul,
		"0.7.5":  EVMVersionIstanbul,
		"0.7.6":  EVMVersionIstanbul,
		"0.8.0":  EVMVersionIstanbul,
		"0.8.1":  EVMVersionIstanbul,
		"0.8.2":  EVMVersionIstanbul,
		"0.8.3":  EVMVersionIstanbul,
		"0.8.4":  EVMVersionIstanbul,
		"0.8.5":  EVMVersionBerlin,
		"0.8.6":  EVMVersionBerlin,
		"0.8.7":  EVMVersionLondon,
		"0.8.8":  EVMVersionLondon,
		"0.8.9":  EVMVersionLondon,
		"0.8.10": EVMVersionLondon,
		"0.8.11": EVMVersionLondon,
		"0.8.12": EVMVersionLondon,
		"0.8.13": EVMVersionLondon,
		"0.8.14": EVMVersionLondon,
		"0.8.15": EVMVersionLondon,
		"0.8.16": EVMVersionLondon,
		"0.8.17": EVMVersionLondon,
		"0.8.18": EVMVersionParis,
		"0.8.19": EVMVersionParis,
		"0.8.20": EVMVersionShanghai,
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
