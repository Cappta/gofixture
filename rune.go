package gofixture

import "unicode/utf8"

// NextValidRune returns the next non-control rune after the current one
func NextValidRune(current rune) rune {
	runeIndex := int(current) + 1
	for utf8ValidRune(rune(runeIndex)) == false {
		runeIndex++
	}
	return rune(runeIndex)
}

// AnyValidRune returns any valid rune
func AnyValidRune() rune {
	value := AnyRune()
	for utf8ValidRune(value) == false {
		value = AnyRune()
	}
	return value
}

// AnyRune returns any rune
func AnyRune() rune {
	return rune(AnyIntBetween(0, utf8.MaxRune))
}
