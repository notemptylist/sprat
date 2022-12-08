package net

import (
	"testing"

	"github.com/notemptylist/sprat/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool(0)
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAdd(t *testing.T) {
	p := NewTxPool(0)
	tx := core.NewTransaction([]byte("Foobar"))
	assert.Nil(t, p.Add(tx))

	assert.Equal(t, p.Len(), 1)
	tx2 := core.NewTransaction([]byte("Foobar"))
	assert.Nil(t, p.Add(tx2))
	assert.Equal(t, p.Len(), 1)
	p.Flush()

	assert.Equal(t, p.Len(), 0)
}
