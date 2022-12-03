package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("Invalid length of bytes")
		panic(msg)
	}

	var val [32]uint8
	for i := 0; i < 32; i++ {
		val[i] = b[i]
	}
	return Hash(val)

}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}
func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
