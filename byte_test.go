package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestByte(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a zero byte", func() {
			zeroByte := byte(0)
			Convey("Then should panic when getting Any Byte Less Than", func() {
				So(func() { AnyByteLessThan(zeroByte) }, ShouldPanic)
			})
			Convey("When getting Any Byte Except", func() {
				anyOtherByte := AnyByteExcept(zeroByte)
				Convey("Then resulting byte should not equal zero byte", func() {
					So(anyOtherByte, ShouldNotEqual, zeroByte)
				})
				Convey("When getting Any Byte Less Than", func() {
					anyByteLessThanAnyOtherByte := AnyByteLessThan(anyOtherByte)
					Convey("Then resulting byte should not equal zero byte", func() {
						So(anyByteLessThanAnyOtherByte, ShouldBeLessThan, anyOtherByte)
					})
				})
			})
		})
	})
}
