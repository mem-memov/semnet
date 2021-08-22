package bit

import (
	"fmt"
)

type Entity struct {
	identifier uint
	storage    storage
}

func newEntity(identifier uint, storage storage) Entity {
	return Entity{
		identifier: identifier,
		storage:    storage,
	}
}

func (e Entity) Identifier() uint {
	return e.identifier
}

func (e Entity) Is(bit bool) bool {
	if bit {
		return e.identifier == bitOneNode
	} else {
		return e.identifier == bitZeroNode
	}
}

func (e Entity) Bit() bool {

	if e.identifier == bitZeroNode {
		return false
	} else {
		return true
	}
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.storage.Connect(sourceIdentifier, e.identifier)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	targets, err := e.storage.ReadTargets(e.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := e.storage.Create()
		if err != nil {
			return 0, err
		}

		err = e.storage.Connect(e.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", e.identifier, len(targets))
	}
}
