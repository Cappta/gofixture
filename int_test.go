package gofixture

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	ConveyWithSeed(time.Now().UTC().UnixNano(), t, func() {
		Convey("Given a random integer lower than 100 and another between 100 and 200", func() {
			min := rand.Intn(100)
			max := rand.Intn(100) + 100
			Convey("When getting any integer between", func() {
				result := AnyIntBetween(min, max)
				Convey("Then result should be greater than or equal to minimum", func() {
					So(result, ShouldBeGreaterThanOrEqualTo, min)
				})
				Convey("Then result should be less than or equal to maximum", func() {
					So(result, ShouldBeLessThanOrEqualTo, max)
				})
			})
		})
		Convey("Given a random integer", func() {
			first := AnyInt()
			Convey("Given another random integer", func() {
				second := AnyInt()
				Convey("Then first integer should not equal second", func() {
					So(first, ShouldNotEqual, second)
				})
			})
		})
	})
}
