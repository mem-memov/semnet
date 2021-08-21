package code

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
)

type layer struct {
	storage  storage
	entities entities
}

func newLayer(storage storage) *layer {
	return &layer{
		storage: storage,
	}
}

func (l *layer) provideRoot(bitEntity bit.Entity) (Entity, error) {

	bitIdentifier, err := bitEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	bitTargets, err := l.storage.ReadTargets(bitIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var codeIdentifier uint
	var characterIdentifier uint

	switch len(bitTargets) {
	case 0:
		codeIdentifier, err = l.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = l.storage.SetReference(bitIdentifier, codeIdentifier)
		if err != nil {
			return Entity{}, err
		}

		characterIdentifier, err = l.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = l.storage.SetReference(codeIdentifier, characterIdentifier)
		if err != nil {
			return Entity{}, err
		}
	case 1:
		if bitTargets[0] != bitEntity.Identifier() {
			return Entity{}, fmt.Errorf("wrong target %d in code layer at bit %d", bitTargets[0], bitIdentifier)
		}

		_, codeIdentifier, err = l.storage.GetReference(bitIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, characterIdentifier, err = l.storage.GetReference(codeIdentifier)
		if err != nil {
			return Entity{}, err
		}
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in code layer at bit %d", len(bitTargets), bitIdentifier)
	}

	return l.entities.create(bitIdentifier, codeIdentifier, characterIdentifier), nil
}
