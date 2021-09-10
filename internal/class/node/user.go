package node

import "fmt"

type User struct {
	identifier uint
	storage    storage
}

func NewUser(identifier uint, storage storage) User {
	return User{
		identifier: identifier,
		storage:    storage,
	}
}

func (u User) Create() (uint, error) {

	identifier, err := u.storage.Create()
	if err != nil {
		return 0, err
	}

	err = u.storage.Connect(identifier, u.identifier)
	if err != nil {
		return 0, err
	}

	return identifier, nil
}

func (u User) Is(identifier uint) (bool, error) {

	targets, err := u.storage.ReadTargets(identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets at user class node %d", identifier)
	}

	return targets[0] == u.identifier, nil
}
