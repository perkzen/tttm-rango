package utils

import "math/rand"

// RandomInt returns a random integer between min and max
//
// - Upper bound is exclusive
//
// - Lower bound is inclusive.
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
