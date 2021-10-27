package remark

import (
	api "github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
)

type Repository struct {
	storage abstract.Storage
	classRepository abstractClass.Repository
	detailRepository abstractDetail.Repository
	factRepository   abstractFact.Repository
}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	detailRepository abstractDetail.Repository,
	factRepository abstractFact.Repository,
) *Repository {

	return &Repository{
		storage:         storage,
		classRepository: classRepository,
		detailRepository: detailRepository,
		factRepository:  factRepository,
	}
}

func (r *Repository) CreateFirstUserRemark(object string, property string) (api.Remark, error) {

	remark, err := createEntity(r.storage)
	if err != nil {
		return nil, err
	}

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	err = remark.PointToClass(class)
	if err != nil {
		return nil, err
	}

	detail, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return nil, err
	}

	err = detail.PointToRemark(remark)
	if err != nil {
		return nil, err
	}

	fact, err := r.factRepository.CreateFirstUserStoryFact()
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(fact)
	if err != nil {
		return nil, err
	}

	err = fact.PointToRemark(remark)
	if err != nil {
		return nil, err
	}

	return remark, nil
}

func (r *Repository) GetRemark(remarkIdentifier uint) (api.Remark, error) {

	return nil, nil
}
