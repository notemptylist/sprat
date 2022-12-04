package types

import "fmt"

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
