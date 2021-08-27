package detail

import (
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

func (s *star) provideBeam(objectPhraseEntity phrase.Entity, propertyPhraseEntity phrase.Entity) (Entity, error) {

	phraseIdentifier, err := objectPhraseEntity.ProvideDetailTarget(propertyPhraseEntity)
	if err != nil {
		return Entity{}, err
	}

	_, remarkIdentifier, err := s.storage.GetReference(phraseIdentifier)

	if remarkIdentifier == 0 {

		remarkIdentifier, err = s.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = s.storage.SetReference(phraseIdentifier, remarkIdentifier)
		if err != nil {
			return Entity{}, err
		}
	}

	return s.entities.create(phraseIdentifier, remarkIdentifier), nil
}
