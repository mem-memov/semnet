package detail

import (
	"github.com/mem-memov/semnet/internal/phrase"
)

type Repository struct {
	entities         *entities
	phraseRepository *phrase.Repository
	star             *star
}

func NewRepository(storage storage, phraseRepository *phrase.Repository) *Repository {
	entities := newEntities(storage, phraseRepository)

	return &Repository{
		entities:         entities,
		phraseRepository: phraseRepository,
		star:             newStar(storage, entities),
	}
}

func (r *Repository) Provide(object string, property string) (Entity, error) {

	objectPhrase, err := r.phraseRepository.Provide(object)
	if err != nil {
		return Entity{}, err
	}

	propertyPhrase, err := r.phraseRepository.Provide(property)
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.star.provideBeam(objectPhrase, propertyPhrase)
	if err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r *Repository) Extract(entity Entity) (string, string, error) {

	return entity.phraseValues()
}

func (r *Repository) Fetch(remarkIdentifier uint) (Entity, error) {

	return r.entities.createWithRemark(remarkIdentifier)
}