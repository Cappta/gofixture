package gofixture

import (
	"testing"
	"time"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	ConveyWithSeed(time.Now().UTC().UnixNano(), t, func() {
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := AnyIntBetween(256, 1024)

			Convey("Given a String", func() {
				value := String(length)
				Convey("Then rune count should equal provided length", func() {
					So(len([]rune(value)), ShouldEqual, length)
				})
				Convey("Given another String", func() {
					anotherValue := String(length)
					Convey("Then results should resemble", func() {
						So(value, ShouldResemble, anotherValue)
					})
				})
			})
			Convey("Given any string", func() {
				first := AnyString(length)
				Convey("Given another string", func() {
					second := AnyString(length)
					Convey("Then first string should not resemble second", func() {
						So(first, ShouldNotResemble, second)
					})
				})
			})
		})
	})
}
