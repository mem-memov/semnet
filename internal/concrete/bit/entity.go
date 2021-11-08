package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
)

type Entity struct {
	value     bool
	class     uint
	character uint
	storage   abstract.Storage
}

var _ abstractBit.Entity = Entity{}

func newEntity(value bool, class uint, character uint, storage abstract.Storage) Entity {
	return Entity{
		value:     value,
		class:     class,
		character: character,
		storage:   storage,
	}
}

func (e Entity) IsBeginningOfCharacters() (bool, error) {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return false, err
	}

	switch len(targets) {

	case 0:

		return false, nil

	case 1:

		backTargets, err := e.storage.ReadTargets(targets[0])
		if err != nil {
			return false, err
		}

		switch len(backTargets) {

		case 0:

			return false, nil

		case 1:

			if backTargets[0] != e.character {
				return false, fmt.Errorf("bit not pointing to itself: %d", e.character)
			}

			return true, nil

		default:

			return false, fmt.Errorf("bit not pointing to itself: %d", e.character)
		}

	default:
		return false, fmt.Errorf("bit has wrong number of target charactersf: %d at %d", len(targets), e.character)
	}
}

func (e Entity) GetCharacter() uint {

	return e.character
}

func (e Entity) Is(bit bool) bool {

	return e.value == bit
}

func (e Entity) Bit() bool {

	return e.value
}

func (e Entity) MarkCharacter(sourceIdentifier uint) error {

	return e.storage.Connect(sourceIdentifier, e.character)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := e.storage.Create()
		if err != nil {
			return 0, err
		}

		err = e.storage.Connect(e.character, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("bit has wrong number of target characters: %d at %d", len(targets), e.character)
	}
}
