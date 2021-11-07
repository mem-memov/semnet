package node

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
)

type Character struct {
	identifier uint
	storage    abstract.Storage
}

func newCharacter(identifier uint, storage abstract.Storage) Character {
	return Character{
		identifier: identifier,
		storage:    storage,
	}
}

func (c Character) IsBeginningOfCharacters() (bool, error) {

	target, err := c.ProvideSingleTarget()
	if err != nil {
		return false, err
	}

	backTargets, err := c.storage.ReadTargets(target)

	switch len(backTargets) {

	case 0:

		return false, nil

	case 1:

		if backTargets[0] != c.identifier {
			return false, fmt.Errorf("character not pointing to itself: %d", c.identifier)
		}

		return true, nil

	default:

		return false, fmt.Errorf("character not pointing to itself: %d", c.identifier)
	}
}

func (c Character) Identifier() uint {
	return c.identifier
}

func (c Character) Mark(sourceIdentifier uint) error {
	return c.storage.Connect(sourceIdentifier, c.identifier)
}

func (c Character) ProvideSingleTarget() (uint, error) {

	targets, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := c.storage.Create()
		if err != nil {
			return 0, err
		}

		err = c.storage.Connect(c.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("bit cluster %d has too many targets: %d", c.identifier, len(targets))
	}
}
