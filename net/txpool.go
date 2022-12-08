package net

import (
	"github.com/notemptylist/sprat/core"
	"github.com/notemptylist/sprat/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool(size int) *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction, size),
	}
}

// Add adds a transaction to the mempool, the caller is responsible for
// checking if it already exists
func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	if p.Has(hash) {
		return nil
	}
	p.transactions[hash] = tx
	return nil
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}
