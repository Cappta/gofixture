package gofixture

// Ints returns an int slice with the specified length
func Ints(length int) (output []int) {
	output = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = i
	}
	return
}

// AnyInts returns an int slice with the specified length and random values
func AnyInts(length int) (output []int) {
	output = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = AnyInt()
	}
	return
}
