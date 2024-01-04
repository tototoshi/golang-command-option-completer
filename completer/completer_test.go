package completer

import (
	"testing"
)

func Test_slices_contains(t *testing.T) {
	a := [][]rune{
		[]rune("abc"),
		[]rune("def"),
	}
	b := []rune("abc")
	if !runesSliceContains(a, b) {
		t.Error("Expected true, got false")
	}
}
