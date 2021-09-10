package node

import "fmt"

type Detail struct {
	identifier uint
	storage    storage
}

func NewDetail(identifier uint, storage storage) Detail {
	return Detail{
		identifier: identifier,
		storage:    storage,
	}
}

func (d Detail) Create() (uint, error) {

	identifier, err := d.storage.Create()
	if err != nil {
		return 0, err
	}

	err = d.storage.Connect(identifier, d.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (d Detail) Is(identifier uint) (bool, error) {

	targets, err := d.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at detail class node %d", identifier)
	}

	return targets[0] == d.identifier, nil
}
