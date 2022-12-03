package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysSignVerifySuccess(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.PublicKey()
	msg := []byte("Testing signing")
	sig, err := privkey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubkey, msg))
}

func TestKeysSignVerifyFail(t *testing.T) {
	privkey := GeneratePrivateKey()
	goodkey := privkey.PublicKey()

	msg := []byte("Testing signing")
	sig, err := privkey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(goodkey, msg))

	badprivkey := GeneratePrivateKey()
	badkey := badprivkey.PublicKey()
	assert.False(t, sig.Verify(badkey, msg))
	assert.False(t, sig.Verify(goodkey, []byte("Bad message")))
}
