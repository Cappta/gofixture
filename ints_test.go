package gofixture

import (
	"testing"
	"time"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInts(t *testing.T) {
	ConveyWithSeed(time.Now().UTC().UnixNano(), t, func() {
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := AnyIntBetween(256, 1024)

			Convey("Given some Ints", func() {
				someInts := Ints(length)
				Convey("Given some more Ints", func() {
					someMoreInts := Ints(length)
					Convey("Then results should resemble", func() {
						So(someInts, ShouldResemble, someMoreInts)
					})
				})
			})
			Convey("Given some random Ints", func() {
				someRandomInts := AnyInts(length)
				Convey("Given some more random Ints", func() {
					someMoreRandomInts := AnyInts(length)
					Convey("Then results should not resemble", func() {
						So(someRandomInts, ShouldNotResemble, someMoreRandomInts)
					})
				})
			})
		})
	})
}
