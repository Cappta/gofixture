package gofixture

import "unicode/utf8"

const (
	firstInvalidRune  = 0
	firstValidRune    = 32
	secondInvalidRune = 127
	secondValidRune   = 160
)

// NextValidRune returns the next non-control rune after the current one
func NextValidRune(current rune) rune {
	runeIndex := current + 1
	for ValidRune(runeIndex) == false {
		runeIndex++
	}
	return runeIndex
}

// AnyValidRune returns any valid rune
func AnyValidRune() rune {
	value := AnyRune()
	for ValidRune(value) == false {
		value = AnyRune()
	}
	return value
}

// AnyRune returns any rune
func AnyRune() rune {
	return rune(AnyIntBetween(0, utf8.MaxRune))
}

// ValidRune returns true if a rune is valid and not a control character
func ValidRune(r rune) bool {
	if r >= firstInvalidRune && r < firstValidRune {
		return false
	} else if r >= secondInvalidRune && r < secondValidRune {
		return false
	}
	return utf8.ValidRune(r)
}
