package numbers

import (
	"math"
	"math/rand"
	"time"
)

// RandomIntBetween RandomIntBetween
func RandomIntBetween(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomInt RandomInt
func RandomInt() int {
	max := int(math.Pow(10, 10))
	return RandomIntBetween(0, max)
}
