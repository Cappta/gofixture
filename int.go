package gofixture

import "math/rand"

// Ints returns an int slice with the specified length
func Ints(length int) (output []int) {
	output = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = i
	}
	return
}

// RandomInts returns an int slice with the specified length and random values
func RandomInts(length int) (output []int) {
	output = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = int(rand.Uint64())
	}
	return
}
