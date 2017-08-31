package gofixture

import (
	"testing"
	"time"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRune(t *testing.T) {
	ConveyWithSeed(time.Now().UTC().UnixNano(), t, func() {
		Convey("Given any rune", func() {
			anyRune := AnyRune()
			Convey("Given another rune", func() {
				anotherRune := AnyRune()
				Convey("Then first rune should not equal second", func() {
					So(anyRune, ShouldNotEqual, anotherRune)
				})
			})
		})
		Convey("Given any valid rune", func() {
			anyValidRune := AnyValidRune()
			Convey("Given another valid rune", func() {
				anotherValidRune := AnyValidRune()
				Convey("Then first rune should not equal second", func() {
					So(anyValidRune, ShouldNotEqual, anotherValidRune)
				})
			})
		})
	})
	Convey("Given the rune 0", t, func() {
		value := rune(0)
		Convey("When retrieving the next valid rune", func() {
			value = NextValidRune(value)
			Convey("Then rune's value should equal 32", func() {
				So(value, ShouldEqual, 32)
			})
		})
	})
	Convey("Given the rune 125", t, func() {
		value := rune(125)
		Convey("When retrieving the next valid rune", func() {
			value = NextValidRune(value)
			Convey("Then rune's value should equal 126", func() {
				So(value, ShouldEqual, 126)
			})
			Convey("When retrieving the next valid rune", func() {
				value = NextValidRune(value)
				Convey("Then rune's value should equal 160", func() {
					So(value, ShouldEqual, 160)
				})
			})
		})
	})
}
