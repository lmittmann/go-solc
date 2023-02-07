//go:build linux

package solc

var solcBaseURL = "https://binaries.soliditylang.org/linux-amd64/"

var solcVersions = map[string]solcVersion{
	"0.8.10": {
		Path:   "solc-linux-amd64-v0.8.10+commit.fc410830",
		Sha256: h("0xc7effacf28b9d64495f81b75228fbf4266ac0ec87e8f1adc489ddd8a4dd06d89"),
	},
	"0.8.11": {
		Path:   "solc-linux-amd64-v0.8.11+commit.d7f03943",
		Sha256: h("0x717c239f3a1dc3a4834c16046a0b4b9f46964665c8ffa82051a6d09fe741cd4f"),
	},
	"0.8.12": {
		Path:   "solc-linux-amd64-v0.8.12+commit.f00d7308",
		Sha256: h("0x556c3ec44faf8ff6b67933fa8a8a403abe82c978d6e581dbfec4bd07360bfbf3"),
	},
	"0.8.13": {
		Path:   "solc-linux-amd64-v0.8.13+commit.abaa5c0e",
		Sha256: h("0xa805dffa86ccd8ed5c9cd18ffcfcca6ff46f635216aa7fc0246546f7be413d62"),
	},
	"0.8.14": {
		Path:   "solc-linux-amd64-v0.8.14+commit.80d49f37",
		Sha256: h("0xd5b027c86c0f8fecc024d5d4f95d8ea48d8a942d79970310e342370532b502f0"),
	},
	"0.8.15": {
		Path:   "solc-linux-amd64-v0.8.15+commit.e14f2714",
		Sha256: h("0x5189155ce322d57fb75e8518d9b39139627edea4fb25b5f0ebed0391c52e74cc"),
	},
	"0.8.16": {
		Path:   "solc-linux-amd64-v0.8.16+commit.07a7930e",
		Sha256: h("0x1632786c6c1f856a4a899232ec975a12f305118f43cce90e724ed0b2eebfeee1"),
	},
	"0.8.17": {
		Path:   "solc-linux-amd64-v0.8.17+commit.8df45f5f",
		Sha256: h("0x99f2070b776e9714f1f76c43c229cf99b8978a92938ee8d2364c6de11c1a03d4"),
	},
	"0.8.18": {
		Path:   "solc-linux-amd64-v0.8.18+commit.87f61d96",
		Sha256: h("0x95e6ed4949a63ad89afb443ecba1fb8302dd2860ee5e9baace3e674a0f48aa77"),
	},
}
