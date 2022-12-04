package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeBlockchain(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestBlockchain(t *testing.T) {
	bc := makeBlockchain(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := makeBlockchain(t)
	assert.True(t, bc.HasBlock(0))

}

func TestAddBlock(t *testing.T) {
	bc := makeBlockchain(t)

	lenBlocks := 100
	for i := 0; i < lenBlocks; i++ {
		block := randomSignedBlock(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.True(t, bc.HasBlock(uint32(lenBlocks)))
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.False(t, bc.HasBlock(uint32(lenBlocks+1)))
	assert.NotNil(t, bc.AddBlock(randomBlock(80)))
}

func TestAddBlockTooHigh(t *testing.T) {
	bc := makeBlockchain(t)
	assert.Equal(t, bc.Height(), uint32(0))
	assert.NotNil(t, bc.AddBlock(randomSignedBlock(t, 3)))

}
