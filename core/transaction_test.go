package core

import (
	"testing"

	"github.com/notemptylist/sprat/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	data := []byte("foo")
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)

}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	data := []byte("foo")
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())
	badKey := crypto.GeneratePrivateKey()
	tx.PublicKey = badKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
