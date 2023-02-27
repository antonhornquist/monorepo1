package uniqueid

import (
	"time"
	"math/rand"
)

func PseudoUniqueId() string {
	var alphanumericChars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	n := 5

	rand.Seed(time.Now().UnixNano())

    b := make([]rune, n)
    for i := range b {
        b[i] = alphanumericChars[rand.Intn(len(alphanumericChars))]
    }
    return string(b)
}

