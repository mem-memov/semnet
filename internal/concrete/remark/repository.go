package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	"github.com/mem-memov/semnet/internal/concrete/detail"
)

type Repository struct {
	factory abstractRemark.Factory
	chain   abstractRemark.Chain
}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, detailRepository *detail.Repository) *Repository {
	factory := NewFactory(storage, classRepository, detailRepository)

	return &Repository{
		factory:         factory,
		chain:            newChain(storage, factory),
	}
}

func (r *Repository) CreateFirstRemark(
	factIdentifier uint,
	object string,
	property string,
) (Remark, error) {

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

func (r *Repository) GetRemark(remarkIdentifier uint) (Remark, error) {

	return nil, nil
}
