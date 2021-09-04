package remark

import (
	"github.com/mem-memov/semnet/internal/detail"
)

type Repository struct {
	entities         *entities
	detailRepository *detail.Repository
	chain            *chain
}

func NewRepository(storage storage, detailRepository *detail.Repository) *Repository {
	entities := newEntities(storage, detailRepository)

	return &Repository{
		entities:         entities,
		detailRepository: detailRepository,
		chain:            newChain(storage, entities),
	}
}

func (r *Repository) AddStoryToUser(userIdentifier uint, remarkEntities []Entity, object string, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return 0, err
	}

	entity, err := r.chain.createFirstLink(detailEntity)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *Repository) AddFactToStory(userIdentifier uint, remarkEntities []Entity, object string, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return 0, err
	}

	entity, err := r.chain.createFirstLink(detailEntity)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *Repository) AddRemarkToFact(userIdentifier uint, remarkEntities []Entity, object string, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return 0, err
	}

	entity, err := r.chain.createFirstLink(detailEntity)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *Repository) Fetch(remarkIdentifier uint) (Entity, error) {

	return Entity{}, nil
}
