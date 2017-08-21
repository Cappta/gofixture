package gofixture

var (
	// StringRunes is a slice containing all the runes available when creating a string
	StringRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890çÇ-=!@#$%¨&*()__+´[~],.;/`{^}<>:?\\|'\"")
)

// String creates a string of the specified length with predictable values from StringRunes sequentially
func String(length int) string {
	data := make([]rune, length)
	stringRunesLength := len(StringRunes)
	for i := 0; i < length; i++ {
		data[i] = StringRunes[i%stringRunesLength]
	}
	return string(data)
}
