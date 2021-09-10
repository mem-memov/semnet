package node

import "fmt"

type Word struct {
	identifier uint
	storage    storage
}

func NewWord(identifier uint, storage storage) Word {
	return Word{
		identifier: identifier,
		storage:    storage,
	}
}

func (w Word) Create() (uint, error) {

	identifier, err := w.storage.Create()
	if err != nil {
		return 0, err
	}

	err = w.storage.Connect(identifier, w.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (w Word) Is(identifier uint) (bool, error) {

	targets, err := w.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at word class node %d", identifier)
	}

	return targets[0] == w.identifier, nil
}
