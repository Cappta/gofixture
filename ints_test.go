package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInts(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
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
