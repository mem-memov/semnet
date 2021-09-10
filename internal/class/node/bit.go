package node

import "fmt"

type Bit struct {
	identifier uint
	storage    storage
}

func NewBit(identifier uint, storage storage) Bit {
	return Bit{
		identifier: identifier,
		storage:    storage,
	}
}

func (b Bit) Create() (uint, error) {

	identifier, err := b.storage.Create()
	if err != nil {
		return 0, err
	}

	err = b.storage.Connect(identifier, b.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (b Bit) Is(identifier uint) (bool, error) {

	targets, err := b.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at bit class node %d", identifier)
	}

	return targets[0] == b.identifier, nil
}

func (b Bit) GetAll() ([]uint, error) {

	return b.storage.ReadSources(b.identifier)
}
