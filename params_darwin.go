//go:build darwin

package solc

var solcBaseURL = "https://binaries.soliditylang.org/macosx-amd64/"

var solcVersions = map[string]solcVersion{
	"0.8.10": {
		Path:   "solc-macosx-amd64-v0.8.10+commit.fc410830",
		Sha256: h("0xa79fff23aeb35be856e446827c44a9cfa4c382f29babd2f6a405ef73d1e2a4cc"),
	},
	"0.8.11": {
		Path:   "solc-macosx-amd64-v0.8.11+commit.d7f03943",
		Sha256: h("0x10cdcc8d8ea4dde9fd8b953b95885dc737f24b8a31fea65f4715ffd007b80281"),
	},
	"0.8.12": {
		Path:   "solc-macosx-amd64-v0.8.12+commit.f00d7308",
		Sha256: h("0x95738a27909a13502385e9fe8f8f3d8a873d2faf5d06ff617bc2fe3edb8c4bf9"),
	},
	"0.8.13": {
		Path:   "solc-macosx-amd64-v0.8.13+commit.abaa5c0e",
		Sha256: h("0x14d4ef013ea82ad95e91fd949b7fa7b78271a483ff1a79c43d6cc58b826f5bea"),
	},
	"0.8.14": {
		Path:   "solc-macosx-amd64-v0.8.14+commit.80d49f37",
		Sha256: h("0xb3d19ab47657af37be4c551f83494248e99d7ba103b6072e8c08dbb62708e2b0"),
	},
	"0.8.15": {
		Path:   "solc-macosx-amd64-v0.8.15+commit.e14f2714",
		Sha256: h("0x00656dc73224e4c0702940df10310bdc024b60f4a7598e774d305bc3b94f7d79"),
	},
	"0.8.16": {
		Path:   "solc-macosx-amd64-v0.8.16+commit.07a7930e",
		Sha256: h("0x7d471cb9bae9a7f29c7ebf402f7e16fa8226b17ba9ab68a88ce107114479dc4d"),
	},
	"0.8.17": {
		Path:   "solc-macosx-amd64-v0.8.17+commit.8df45f5f",
		Sha256: h("0xe40eef83c24d4c42b47f461b01748a6ca89f1e09e778995b71debfa0de99e12a"),
	},
}
