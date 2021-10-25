package detail

import (
	"github.com/mem-memov/semnet/internal/abstract/detail"
	"github.com/mem-memov/semnet/internal/abstract/phrase"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Repository struct {
	entities         *entities
	phraseRepository phrase.Repository
	star             *star
}

var _ detail.Repository = &Repository{}

func NewRepository(storage storage, classRepository *class.Repository, phraseRepository phrase.Repository) *Repository {
	entities := newEntities(storage, classRepository, phraseRepository)

	return &Repository{
		entities:         entities,
		phraseRepository: phraseRepository,
		star:             newStar(storage, entities),
	}
}

func (r *Repository) Extend(objectIdentifier uint, property string) (detail.Entity, error) {

	objectPhrase, err := r.phraseRepository.Fetch(objectIdentifier)
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

func (r *Repository) Provide(object string, property string) (detail.Entity, error) {

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

func (r *Repository) Fetch(remarkIdentifier uint) (detail.Entity, error) {

	return r.entities.createWithRemark(remarkIdentifier)
}
