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

func (e Entity) GetCharacter() uint {

	return e.character
}

func (e Entity) HasTargetCharacter() (bool, error) {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return false, err
	}

	switch len(targets) {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("bit has wrong number of target characters: %d at %d", len(targets), e.character)
	}
}

func (e Entity) PointToCharacter(character uint) error {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return err
	}

	if len(targets) != 0 {
		return fmt.Errorf("target character cannot be added to bit: %d", e.character)
	}

	return e.storage.Connect(e.character, character)
}

func (e Entity) GetTargetCharacter() (uint, error) {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("bit has wrong number of target characters: %d at %d", len(targets), e.character)
	}

	return targets[0], nil
}

func (e Entity) Is(bit bool) bool {

	return e.value == bit
}

func (e Entity) Bit() bool {

	return e.value
}
