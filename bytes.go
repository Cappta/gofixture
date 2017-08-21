package gofixture

// Bytes creates a byte slice of the specified length with predictable values {0, ..., 255, 0, ...}
func Bytes(length int) []byte {
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		data[i] = byte(i)
	}
	return data
}

// ZeroBytes creates a byte slice of the specified length filled with zeroes
func ZeroBytes(length int) []byte {
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		data[i] = 0
	}
	return data
}
