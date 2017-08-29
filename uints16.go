package gofixture

import "math/rand"

// AnyUints16 returns an int16 slice with the specified length and random values
func AnyUints16(length int) (output []uint16) {
	output = make([]uint16, length)
	for i := 0; i < length; i++ {
		output[i] = uint16(rand.Uint32())
	}
	return
}
