package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBytes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a length between 256 and 1024", func() {
			length := AnyIntBetween(256, 1024)

			Convey("Given some Bytes", func() {
				someBytes := Bytes(length)
				Convey("Given some more Bytes", func() {
					someMoreBytes := Bytes(length)
					Convey("Then results should resemble", func() {
						So(someBytes, ShouldResemble, someMoreBytes)
					})
				})
			})
			Convey("Given any Bytes", func() {
				anyBytes := AnyBytes(length)
				Convey("Given more Bytes", func() {
					moreBytes := AnyBytes(length)
					Convey("Then results should not resemble", func() {
						So(anyBytes, ShouldNotResemble, moreBytes)
					})
				})
			})
			Convey("Given some Zero Bytes", func() {
				zeroBytes := ZeroBytes(length)
				Convey("Then each byte should equal zero", func() {
					for i := 0; i < len(zeroBytes); i++ {
						So(zeroBytes[i], ShouldEqual, 0)
					}
				})
			})
		})
	})
}
