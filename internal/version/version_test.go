package version_test

import (
	"strconv"
	"testing"

	"github.com/lmittmann/go-solc/internal/version"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		VersionStr string
		Want       bool
	}{
		{"0.1.2", true},
		{"go0.1.2", false},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := version.IsValid(test.VersionStr)
			if test.Want != got {
				t.Fatalf("want %t, got %t", test.Want, got)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		X, Y string
		Want int
	}{
		{"0.5", "0.5", 0},
		{"0.1", "0.5", -1},
		{"0.5", "0.1", 1},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := version.Compare(test.X, test.Y)
			if test.Want != got {
				t.Fatalf("want %d, got %d", test.Want, got)
			}
		})
	}
}
