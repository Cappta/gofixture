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
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := int(256 + rand.Int31()%(1024-256))

			Convey("Given some Bytes", func() {
				someBytes := Bytes(length)
				Convey("Given some more Bytes", func() {
					someMoreBytes := Bytes(length)
					Convey("Then results should resemble", func() {
						So(someBytes, ShouldResemble, someMoreBytes)
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
