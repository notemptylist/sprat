package core

import (
	"testing"

	"github.com/notemptylist/sprat/types"
	"github.com/stretchr/testify/assert"
)

func makeBlockchain(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
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
		height := uint32(i + 1)
		block := randomSignedBlock(t, height, getPrevBlockHash(t, bc, height))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(80, types.Hash{})))
}

func TestAddBlockTooHigh(t *testing.T) {
	bc := makeBlockchain(t)
	assert.Equal(t, bc.Height(), uint32(0))
	assert.NotNil(t, bc.AddBlock(randomSignedBlock(t, 3, types.Hash{})))

}

func TestGetHeader(t *testing.T) {
	bc := makeBlockchain(t)

	lenBlocks := 100
	for i := 0; i < lenBlocks; i++ {
		height := uint32(i + 1)
		block := randomSignedBlock(t, height, getPrevBlockHash(t, bc, height))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}

}
