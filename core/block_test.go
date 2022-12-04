package core

import (
	"testing"
	"time"

	"github.com/notemptylist/sprat/crypto"
	"github.com/notemptylist/sprat/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("random"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomSignedBlock(t *testing.T, height uint32) *Block {

	privkey := crypto.GeneratePrivateKey()

	b := randomBlock(height)
	b.Sign(privkey)
	assert.Nil(t, b.Sign(privkey))
	return b
}
func TestSignBlock(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privkey))
	assert.NotNil(t, b.Signature)
}
func TestVerifyBlock(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := randomBlock(101)
	assert.Nil(t, b.Sign(privkey))
	assert.Equal(t, b.Validator, privkey.PublicKey())
	assert.Nil(t, b.Verify())

	otherkey := crypto.GeneratePrivateKey()
	b.Validator = otherkey.PublicKey()
	assert.NotNil(t, b.Verify())

	// b.Height = 100
	// assert.Nil(t, b.Verify())
}
