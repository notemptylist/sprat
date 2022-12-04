package core

import "fmt"

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)
	return bc, bc.addBlockWithoutValidation(genesis)
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) Height() uint32 {
	//validate
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}
	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	fmt.Printf("Added block with hash{%s} height{%d}\n", b.Hash(BlockHasher{}), b.Height)
	return bc.store.Put(b)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {

	if height > bc.Height() {
		return nil, fmt.Errorf("Requested header height (%d) too high", height)
	}
	return bc.headers[height], nil
}
