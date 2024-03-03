package version

import "go/version"

func IsValid(x string) bool {
	return version.IsValid("go" + x)
}

func Compare(x, y string) int {
	return version.Compare("go"+x, "go"+y)
}
