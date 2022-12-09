package net

import (
	"math/rand"
	"strconv"
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

func TestSortTransactions(t *testing.T) {
	txLen := 10
	p := NewTxPool(txLen)

	for i := 0; i < txLen; i++ {
		// create a transaction with unique data
		tx := core.NewTransaction([]byte("Foobar" + strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(int64(i + rand.Intn(1000)))
		assert.Nil(t, p.Add(tx))
	}
	assert.Equal(t, txLen, p.Len())

	txx := p.Transactions()
	for i := 0; i < len(txx)-1; i++ {
		assert.True(t, txx[i].GetFirstSeen() < txx[i+1].GetFirstSeen())
	}
}
