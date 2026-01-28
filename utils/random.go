package utils

import (
	"fmt"
	"math/rand" // Use math/rand for general purpose random data
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Seed the random number generator so results change every time the app runs
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ { // FIXED: Was 1 < n, which caused an infinite loop
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurency() string {
	currencies := []string{"EUR", "USD", "CFA"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a randome email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
