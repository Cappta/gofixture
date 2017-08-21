package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := int(256 + rand.Int31()%(1024-256))

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
				someRandomInts := RandomInts(length)
				Convey("Given some more random Ints", func() {
					someMoreRandomInts := RandomInts(length)
					Convey("Then results should not resemble", func() {
						So(someRandomInts, ShouldNotResemble, someMoreRandomInts)
					})
				})
			})
		})
	})
}
