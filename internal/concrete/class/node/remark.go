package node

import "fmt"

type Remark struct {
	identifier uint
	storage    storage
}

func NewRemark(identifier uint, storage storage) Remark {
	return Remark{
		identifier: identifier,
		storage:    storage,
	}
}

func (r Remark) Create() (uint, error) {

	identifier, err := r.storage.Create()
	if err != nil {
		return 0, err
	}

	err = r.storage.Connect(identifier, r.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (r Remark) Is(identifier uint) (bool, error) {

	targets, err := r.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at position class node %d", identifier)
	}

	return targets[0] == r.identifier, nil
}
