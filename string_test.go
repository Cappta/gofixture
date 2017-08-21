package gofixture

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a length higher than 256 and lower than 1024", func() {
			length := int(256 + rand.Int31()%(1024-256))

			Convey("Given a String", func() {
				string := String(length)
				Convey("Given another String", func() {
					anotherString := String(length)
					Convey("Then results should resemble", func() {
						So(string, ShouldResemble, anotherString)
					})
				})
			})
		})
	})
}
