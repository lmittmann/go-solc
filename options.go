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
	defaultEVMVersions = map[SolcVersion]EVMVersion{
		SolcVersion0_5_0:  EVMVersionByzantium,
		SolcVersion0_5_1:  EVMVersionByzantium,
		SolcVersion0_5_2:  EVMVersionByzantium,
		SolcVersion0_5_3:  EVMVersionByzantium,
		SolcVersion0_5_4:  EVMVersionByzantium,
		SolcVersion0_5_5:  EVMVersionPetersburg,
		SolcVersion0_5_6:  EVMVersionPetersburg,
		SolcVersion0_5_7:  EVMVersionPetersburg,
		SolcVersion0_5_8:  EVMVersionPetersburg,
		SolcVersion0_5_9:  EVMVersionPetersburg,
		SolcVersion0_5_10: EVMVersionPetersburg,
		SolcVersion0_5_11: EVMVersionPetersburg,
		SolcVersion0_5_12: EVMVersionPetersburg,
		SolcVersion0_5_13: EVMVersionPetersburg,
		SolcVersion0_5_14: EVMVersionIstanbul,
		SolcVersion0_5_15: EVMVersionIstanbul,
		SolcVersion0_5_16: EVMVersionIstanbul,
		SolcVersion0_5_17: EVMVersionIstanbul,
		SolcVersion0_6_0:  EVMVersionIstanbul,
		SolcVersion0_6_1:  EVMVersionIstanbul,
		SolcVersion0_6_2:  EVMVersionIstanbul,
		SolcVersion0_6_3:  EVMVersionIstanbul,
		SolcVersion0_6_4:  EVMVersionIstanbul,
		SolcVersion0_6_5:  EVMVersionIstanbul,
		SolcVersion0_6_6:  EVMVersionIstanbul,
		SolcVersion0_6_7:  EVMVersionIstanbul,
		SolcVersion0_6_8:  EVMVersionIstanbul,
		SolcVersion0_6_9:  EVMVersionIstanbul,
		SolcVersion0_6_10: EVMVersionIstanbul,
		SolcVersion0_6_11: EVMVersionIstanbul,
		SolcVersion0_6_12: EVMVersionIstanbul,
		SolcVersion0_7_0:  EVMVersionIstanbul,
		SolcVersion0_7_1:  EVMVersionIstanbul,
		SolcVersion0_7_2:  EVMVersionIstanbul,
		SolcVersion0_7_3:  EVMVersionIstanbul,
		SolcVersion0_7_4:  EVMVersionIstanbul,
		SolcVersion0_7_5:  EVMVersionIstanbul,
		SolcVersion0_7_6:  EVMVersionIstanbul,
		SolcVersion0_8_0:  EVMVersionIstanbul,
		SolcVersion0_8_1:  EVMVersionIstanbul,
		SolcVersion0_8_2:  EVMVersionIstanbul,
		SolcVersion0_8_3:  EVMVersionIstanbul,
		SolcVersion0_8_4:  EVMVersionIstanbul,
		SolcVersion0_8_5:  EVMVersionBerlin,
		SolcVersion0_8_6:  EVMVersionBerlin,
		SolcVersion0_8_7:  EVMVersionLondon,
		SolcVersion0_8_8:  EVMVersionLondon,
		SolcVersion0_8_9:  EVMVersionLondon,
		SolcVersion0_8_10: EVMVersionLondon,
		SolcVersion0_8_11: EVMVersionLondon,
		SolcVersion0_8_12: EVMVersionLondon,
		SolcVersion0_8_13: EVMVersionLondon,
		SolcVersion0_8_14: EVMVersionLondon,
		SolcVersion0_8_15: EVMVersionLondon,
		SolcVersion0_8_16: EVMVersionLondon,
		SolcVersion0_8_17: EVMVersionLondon,
		SolcVersion0_8_18: EVMVersionParis,
		SolcVersion0_8_19: EVMVersionParis,
		SolcVersion0_8_20: EVMVersionShanghai,
		SolcVersion0_8_21: EVMVersionShanghai,
		SolcVersion0_8_22: EVMVersionShanghai,
		SolcVersion0_8_23: EVMVersionShanghai,
		SolcVersion0_8_24: EVMVersionShanghai,
		SolcVersion0_8_25: EVMVersionCancun,
		SolcVersion0_8_26: EVMVersionCancun,
		SolcVersion0_8_27: EVMVersionCancun,
		SolcVersion0_8_28: EVMVersionCancun,
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
