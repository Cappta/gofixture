package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRune(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given any valid rune", func() {
			validRune := AnyValidRune()
			Convey("Given another valid rune", func() {
				anotherValidRune := AnyValidRune()
				Convey("Then first rune should not equal second", func() {
					So(validRune, ShouldNotEqual, anotherValidRune)
				})
			})
			Convey("When the next rune is invalid", func() {
				utf8ValidRune = func(r rune) bool {
					if r == rune(int(validRune)+1) {
						return false
					}
					return true
				}
				Convey("When retrieving the next valid rune", func() {
					nextValidRune := NextValidRune(validRune)
					Convey("Then the next valid rune should equal the second next rune", func() {
						So(nextValidRune, ShouldEqual, rune(int(validRune)+2))
					})
				})
			})
		})
		Convey("When the next rune is invalid", func() {
			counter := 0
			var firstRune rune
			utf8ValidRune = func(r rune) bool {
				counter++
				if counter == 1 {
					firstRune = r
					return false
				}
				return true
			}
			Convey("When retrieving any valid rune", func() {
				finalRune := AnyValidRune()
				Convey("Then the first attempted rune should not equal the final rune", func() {
					So(firstRune, ShouldNotEqual, finalRune)
				})
				Convey("Then the rune check counter should equal 2", func() {
					So(counter, ShouldEqual, 2)
				})
			})
		})
	})
}
