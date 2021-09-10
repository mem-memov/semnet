package node

import "fmt"

type Phrase struct {
	identifier uint
	storage    storage
}

func NewPhrase(identifier uint, storage storage) Phrase {
	return Phrase{
		identifier: identifier,
		storage:    storage,
	}
}

func (p Phrase) Create() (uint, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return 0, err
	}

	err = p.storage.Connect(identifier, p.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (p Phrase) Is(identifier uint) (bool, error) {

	targets, err := p.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at phrase class node %d", identifier)
	}

	return targets[0] == p.identifier, nil
}
