package uniqueid

import (
    "testing"
)

func TestLengthOfPseudoUniqueId(t *testing.T) {
	expected := 6
	actual := len(PseudoUniqueId())
    if actual != expected {
        t.Fatalf(`len(PseudoUniqueId()) = %d, expected %d`, actual, expected)
    }
}
