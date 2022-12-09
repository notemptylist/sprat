package net

import (
	"sort"

	"github.com/notemptylist/sprat/core"
	"github.com/notemptylist/sprat/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

type TxSorter struct {
	txx []*core.Transaction
}

func NewTxSorter(txMap map[types.Hash]*core.Transaction) *TxSorter {
	txx := make([]*core.Transaction, len(txMap))

	i := 0
	for _, val := range txMap {
		txx[i] = val
		i++
	}
	s := &TxSorter{txx}
	sort.Sort(s)
	return s
}

func (tx *TxSorter) Len() int { return len(tx.txx) }
func (tx *TxSorter) Swap(i, j int) {
	tx.txx[i], tx.txx[j] = tx.txx[j], tx.txx[i]
}
func (tx *TxSorter) Less(i, j int) bool {
	return tx.txx[i].GetFirstSeen() < tx.txx[j].GetFirstSeen()
}

func NewTxPool(size int) *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction, size),
	}
}

func (p *TxPool) Transactions() []*core.Transaction {
	sorter := NewTxSorter(p.transactions)
	return sorter.txx
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
