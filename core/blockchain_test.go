package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)

	fmt.Println(bc.Height())
}
