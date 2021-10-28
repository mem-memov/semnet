package node

import "fmt"

type Fact struct {
	identifier uint
	storage    storage
}

func NewFact(identifier uint, storage storage) Fact {
	return Fact{
		identifier: identifier,
		storage:    storage,
	}
}

func (f Fact) GetIdentifier() uint {
	return f.identifier
}

func (f Fact) Create() (uint, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return 0, err
	}

	err = f.storage.Connect(identifier, f.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (f Fact) Is(identifier uint) (bool, error) {

	targets, err := f.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at fact class node %d", identifier)
	}

	return targets[0] == f.identifier, nil
}
