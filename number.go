package armory

import (
	"math"
	"math/rand"
	"time"
)

type number struct{}

var Number *number

// RandomIntBetween RandomIntBetween
func (n *number) RandomIntBetween(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomInt RandomInt
func (n *number) RandomInt() int {
	max := int(math.Pow(10, 10))
	return n.RandomIntBetween(0, max)
}
