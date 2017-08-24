package gofixture

import "math/rand"

// AnyIntBetween returns a random positive integer between minimum and maximum values
func AnyIntBetween(min, max int) int {
	return min + int(rand.Int31())%(max-min)
}
