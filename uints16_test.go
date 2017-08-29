package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUints16(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := AnyIntBetween(256, 1024)

			Convey("Given some random Ints", func() {
				someRandomInts := AnyUints16(length)
				Convey("Given some more random Ints", func() {
					someMoreRandomInts := AnyUints16(length)
					Convey("Then results should not resemble", func() {
						So(someRandomInts, ShouldNotResemble, someMoreRandomInts)
					})
				})
			})
		})
	})
}
