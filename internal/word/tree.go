package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/character"
)

type tree struct {
	storage  storage
	entities *entities
}

func newLayer(storage storage, entities *entities) *tree {
	return &tree{
		storage:  storage,
		entities: entities,
	}
}

func (t *tree) provideRoot(characterEntity character.Entity) (Entity, error) {

	bitIdentifier, err := characterEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	bitTargets, err := t.storage.ReadTargets(bitIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var characterIdentifier uint
	var characterIdentifier uint

	switch len(bitTargets) {
	case 0:
		err = characterEntity.Mark(bitIdentifier)
		if err != nil {
			return Entity{}, err
		}

		characterIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(bitIdentifier, characterIdentifier)
		if err != nil {
			return Entity{}, err
		}

		characterIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(characterIdentifier, characterIdentifier)
		if err != nil {
			return Entity{}, err
		}
	case 1:
		if bitTargets[0] != characterEntity.Identifier() {
			return Entity{}, fmt.Errorf("wrong target %d in character tree at bit %d", bitTargets[0], bitIdentifier)
		}

		_, characterIdentifier, err = t.storage.GetReference(bitIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, characterIdentifier, err = t.storage.GetReference(characterIdentifier)
		if err != nil {
			return Entity{}, err
		}
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in character tree at bit %d", len(bitTargets), bitIdentifier)
	}

	return t.entities.create(bitIdentifier, characterIdentifier, characterIdentifier), nil
}
