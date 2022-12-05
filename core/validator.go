package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("Chain already has block (%d) with has (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("Block (%s) too high (%d)", b.Hash(BlockHasher{}), b.Height)
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if b.PrevBlockHash != hash {
		return fmt.Errorf("prev header hash doesn't match")
	}

	// Verify the block and all of its transactions
	if err := b.Verify(); err != nil {
		return err
	}
	return nil
}
