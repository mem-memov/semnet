package code

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
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

func (t *tree) provideRoot(bitEntity bit.Entity) (Entity, error) {

	bitIdentifier, err := bitEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	bitTargets, err := t.storage.ReadTargets(bitIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var codeIdentifier uint
	var characterIdentifier uint

	switch len(bitTargets) {
	case 0:
		err = bitEntity.Mark(bitIdentifier)
		if err != nil {
			return Entity{}, err
		}

		codeIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(bitIdentifier, codeIdentifier)
		if err != nil {
			return Entity{}, err
		}

		characterIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(codeIdentifier, characterIdentifier)
		if err != nil {
			return Entity{}, err
		}
	case 1:
		if bitTargets[0] != bitEntity.Identifier() {
			return Entity{}, fmt.Errorf("wrong target %d in code tree at bit %d", bitTargets[0], bitIdentifier)
		}

		_, codeIdentifier, err = t.storage.GetReference(bitIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, characterIdentifier, err = t.storage.GetReference(codeIdentifier)
		if err != nil {
			return Entity{}, err
		}
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in code tree at bit %d", len(bitTargets), bitIdentifier)
	}

	return t.entities.create(bitIdentifier, codeIdentifier, characterIdentifier), nil
}
