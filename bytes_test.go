package gofixture

import (
	"testing"
	"time"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBytes(t *testing.T) {
	ConveyWithSeed(time.Now().UTC().UnixNano(), t, func() {
		Convey("Given a length between 256 and 1024", func() {
			length := AnyIntBetween(256, 1024)

			Convey("Given some Bytes", func() {
				someBytes := Bytes(length)
				Convey("Then bytes length should equal provided length", func() {
					So(len(someBytes), ShouldEqual, length)
				})
				Convey("Given some more Bytes", func() {
					someMoreBytes := Bytes(length)
					Convey("Then results should resemble", func() {
						So(someBytes, ShouldResemble, someMoreBytes)
					})
				})
			})
			Convey("Given any Bytes", func() {
				anyBytes := AnyBytes(length)
				Convey("Then bytes length should equal provided length", func() {
					So(len(anyBytes), ShouldEqual, length)
				})
				Convey("Given more Bytes", func() {
					moreBytes := AnyBytes(length)
					Convey("Then results should not resemble", func() {
						So(anyBytes, ShouldNotResemble, moreBytes)
					})
				})
			})
			Convey("Given some Zero Bytes", func() {
				zeroBytes := ZeroBytes(length)
				Convey("Then bytes length should equal provided length", func() {
					So(len(zeroBytes), ShouldEqual, length)
				})
				Convey("Then each byte should equal zero", func() {
					for i := 0; i < len(zeroBytes); i++ {
						So(zeroBytes[i], ShouldEqual, 0)
					}
				})
			})
		})
	})
}
