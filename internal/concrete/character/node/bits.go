package node

import (
	"github.com/mem-memov/semnet/internal/concrete/bit"
)

type Bits struct {
	storage       storage
	bitRepository *bit.Repository
}

func NewBits(storage storage, bitRepository *bit.Repository) *Bits {
	return &Bits{
		storage:       storage,
		bitRepository: bitRepository,
	}
}

func (b *Bits) Create(identifier uint) Bit {
	return newBit(identifier, b.storage, b.bitRepository)
}
