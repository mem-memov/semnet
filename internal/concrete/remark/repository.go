package remark

import (
	api "github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	"github.com/mem-memov/semnet/internal/concrete/detail"
)

type Repository struct {
	factory abstractRemark.Factory
	factRepository   abstractFact.Repository
}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	detailRepository *detail.Repository,
	factRepository abstractFact.Repository,
) *Repository {

	return &Repository{
		factory:         newFactory(storage, classRepository, detailRepository),
		factRepository:  factRepository,
	}
}

func (r *Repository) CreateFirstUserRemark(object string, property string) (api.Remark, error) {



	return Entity{}, nil
}

func (r *Repository) GetRemark(remarkIdentifier uint) (api.Remark, error) {

	return nil, nil
}
