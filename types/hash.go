package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}

	return true
}

func (h Hash) toSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}
	return b
}

func (h Hash) String() string {
	return hex.EncodeToString(h.toSlice())
}

func HashFromBytes(b []byte) Hash {

	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}
	return Hash(value)
}

func RandomBytes(size int) []byte {
	// make() is a built-in function that returns a slice of the given type, length and capacity.
	token := make([]byte, size)
	//rand is a package in Go standard library.
	/*
		rand.read() is a function that reads random
		 bytes from the system's random number
		 generator and writes them into the given slice.
	*/
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
