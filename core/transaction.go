package core

import (
	"fmt"
	"io"

	"github.com/notemptylist/sprat/crypto"
	"github.com/notemptylist/sprat/types"
)

type Transaction struct {
	Data []byte // Arbitrary data

	Sender    crypto.PublicKey
	Signature *crypto.Signature

	hash types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.Sender = privKey.PublicKey()
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Hash(h Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = h.Hash(tx)
	}
	return tx.hash
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("tx has no signature")
	}

	if tx.Signature.Verify(tx.Sender, tx.Data) == false {
		return fmt.Errorf("Invalid transaction signature")
	}
	return nil
}
func (tx *Transaction) DecodeBinary(r io.Reader) error { return nil }
func (tx *Transaction) EncodeBinary(w io.Writer) error { return nil }
