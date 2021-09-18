package remark

import (
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/detail"
)

type Repository struct {
	entities         *entities
	detailRepository *detail.Repository
	chain            *chain
}

func NewRepository(storage storage, classRepository *class.Repository, detailRepository *detail.Repository) *Repository {
	entities := newEntities(storage, detailRepository)

	return &Repository{
		entities:         entities,
		detailRepository: detailRepository,
		chain:            newChain(storage, entities),
	}
}

func (r *Repository) CreateRemark(
	factIdentifier uint,
	remarkEntities []Entity,
	object string,
	property string,
	) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return Entity{}, err
	}

	_, err = r.chain.createFirstLink(detailEntity)
	if err != nil {
		return Entity{}, err
	}

	return Entity{}, nil
}

func (r *Repository) GetRemark(remarkIdentifier uint) (Entity, error) {

	return Entity{}, nil
}
