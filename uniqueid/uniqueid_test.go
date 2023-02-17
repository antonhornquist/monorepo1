package uniqueid

import (
    "testing"
)

func TestLengthOfPseudoUniqueId(t *testing.T) {
	expected := 5
	actual := len(PseudoUniqueId())
    if actual != expected {
        t.Fatalf(`len(PseudoUniqueId()) = %d, expected %d`, actual, expected)
    }
}
