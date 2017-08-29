package gofixture

// String creates a string of the specified length with predictable values from StringRunes sequentially
func String(length int) string {
	data := make([]rune, length)
	var lastRune rune
	for i := 0; i < length; i++ {
		lastRune = NextValidRune(lastRune)
		data[i] = lastRune
	}
	return string(data)
}

// AnyString returns an string with the specified length
func AnyString(length int) string {
	data := make([]rune, length)
	for i := 0; i < length; i++ {
		data[i] = AnyValidRune()
	}
	return string(data)
}
