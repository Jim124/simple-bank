package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)

}

// RandomInt generates a random integer min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1) // min -> max
}

// RandomString generate a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for range n {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
