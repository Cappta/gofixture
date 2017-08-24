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
	})
}
