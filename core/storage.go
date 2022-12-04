package core

type Storage interface {
	Put(*Block) error
}

type MemoryStore struct {
}

func (m *MemoryStore) Put(b *Block) error {
	return nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}
