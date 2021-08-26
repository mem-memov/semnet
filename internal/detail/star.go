package detail

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/phrase"
)

type star struct {
	storage  storage
	entities *entities
}

func newStar(storage storage, entities *entities) *star {
	return &star{
		storage:  storage,
		entities: entities,
	}
}

func (s *star) provideRoot(phraseEntity phrase.Entity) (Entity, error) {

	phraseIdentifier, err := phraseEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	phraseTargets, err := s.storage.ReadTargets(phraseIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var detailIdentifier uint
	var remarkIdentifier uint

	switch len(phraseTargets) {
	case 0:
		err = phraseEntity.Mark(phraseIdentifier)
		if err != nil {
			return Entity{}, err
		}

		detailIdentifier, err = s.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = s.storage.SetReference(phraseIdentifier, detailIdentifier)
		if err != nil {
			return Entity{}, err
		}

		remarkIdentifier, err = s.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = s.storage.SetReference(detailIdentifier, remarkIdentifier)
		if err != nil {
			return Entity{}, err
		}
	case 1:
		if phraseTargets[0] != phraseEntity.DetailIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in detail tree at phrase %d", phraseTargets[0], phraseIdentifier)
		}

		_, phraseIdentifier, err = s.storage.GetReference(phraseIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, detailIdentifier, err = s.storage.GetReference(phraseIdentifier)
		if err != nil {
			return Entity{}, err
		}
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in detail tree at phrase %d", len(phraseTargets), phraseIdentifier)
	}

	return s.entities.create(phraseIdentifier, detailIdentifier, remarkIdentifier), nil
}
