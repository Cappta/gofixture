package gofixture

import "math/rand"

// AnyByteExcept returns any byte except the provided one
func AnyByteExcept(current byte) (output byte) {
	for {
		output = AnyByte()
		if output != current {
			return
		}
	}
}

// AnyByteLessThan returns any byte less than the provided one
func AnyByteLessThan(limit byte) (output byte) {
	if limit == 0 {
		panic("Cannot Fixture to any byte less than 0")
	}
	return byte(rand.Uint32() % uint32(limit))
}

// AnyByte returns any byte
func AnyByte() byte {
	return byte(rand.Uint32())
}
