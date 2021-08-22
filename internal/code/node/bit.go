package node

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
)

type Bit struct {
	identifier    uint
	storage       storage
	bitRepository *bit.Repository
}

func newBit(identifier uint, storage storage, bitRepository *bit.Repository) Bit {
	return Bit{
		identifier:    identifier,
		storage:       storage,
		bitRepository: bitRepository,
	}
}

func (b Bit) Identifier() uint {
	return b.identifier
}

func (b Bit) HasBitValue(value bool) (bool, error) {

	targets, err := b.storage.ReadTargets(b.identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets %d in code layer at bit %d", len(targets), b.identifier)
	}

	bitEntity, err := b.bitRepository.Fetch(targets[0])
	if err != nil {
		return false, err
	}

	return bitEntity.Is(value), nil
}

func (b Bit) NewBit(bitValue bool) (Bit, error) {

	identifier, err := b.storage.Create()
	if err != nil {
		return Bit{}, err
	}

	bitEntity, err := b.bitRepository.Provide(bitValue)
	if err != nil {
		return Bit{}, err
	}

	err = bitEntity.Mark(identifier)
	if err != nil {
		return Bit{}, err
	}

	return newBit(identifier, b.storage, b.bitRepository), nil
}

func (b Bit) BitValue() (bool, error) {

	targets, err := b.storage.ReadTargets(b.identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets %d in code layer at bit %d", len(targets), b.identifier)
	}

	bitEntity, err := b.bitRepository.Fetch(targets[0])
	if err != nil {
		return false, err
	}

	return bitEntity.Bit(), nil
}
