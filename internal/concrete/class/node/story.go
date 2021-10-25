package node

import "fmt"

type Story struct {
	identifier uint
	storage    storage
}

func NewStory(identifier uint, storage storage) Story {
	return Story{
		identifier: identifier,
		storage:    storage,
	}
}

func (s Story) Create() (uint, error) {

	identifier, err := s.storage.Create()
	if err != nil {
		return 0, err
	}

	err = s.storage.Connect(identifier, s.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (s Story) Is(identifier uint) (bool, error) {

	targets, err := s.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at story class node %d", identifier)
	}

	return targets[0] == s.identifier, nil
}
