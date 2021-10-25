package node

import "fmt"

type Character struct {
	identifier uint
	storage    storage
}

func NewCharacter(identifier uint, storage storage) Character {
	return Character{
		identifier: identifier,
		storage:    storage,
	}
}

func (c Character) Create() (uint, error) {

	identifier, err := c.storage.Create()
	if err != nil {
		return 0, err
	}

	err = c.storage.Connect(identifier, c.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (c Character) Is(identifier uint) (bool, error) {

	targets, err := c.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at character class node %d", identifier)
	}

	return targets[0] == c.identifier, nil
}
