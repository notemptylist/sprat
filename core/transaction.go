package core

import (
	"fmt"
	"io"

	"github.com/notemptylist/sprat/crypto"
)

type Transaction struct {
	Data []byte // Arbitrary data

	Sender    crypto.PublicKey
	Signature *crypto.Signature
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
