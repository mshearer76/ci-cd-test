package foo

import "testing"

func TestBar(t *testing.T) {
	t.Logf("Running Bar test\n")
	Bar()
}