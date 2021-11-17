package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Repository struct {
	remarkStorage    abstractRemark.Storage
	classRepository  abstractClass.Repository
	detailRepository abstractDetail.Repository
	factRepository   abstractFact.Repository
}

var _ abstractRemark.Repository = &Repository{}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	detailRepository abstractDetail.Repository,
	factRepository abstractFact.Repository,
) *Repository {

	return &Repository{
		remarkStorage:    NewStorage(storage),
		classRepository:  classRepository,
		detailRepository: detailRepository,
		factRepository:   factRepository,
	}
}

func (r *Repository) CreateFirstUserRemark(object string, property string) (abstractRemark.Aggregate, error) {

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return nil, err
	}

	remark, err := r.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	// detail

	detail, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return nil, err
	}

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return nil, err
	}

	// fact

	fact, err := r.factRepository.CreateFirstStoryFact()
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(fact.GetRemark())
	if err != nil {
		return nil, err
	}

	err = fact.PointToRemark(remark.GetFact())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    r.remarkStorage,
		classRepository:  r.classRepository,
		detailRepository: r.detailRepository,
		factRepository:   r.factRepository,
	}, nil
}

func (r *Repository) GetRemark(remarkIdentifier uint) (abstractRemark.Aggregate, error) {

	remark, err := r.remarkStorage.ReadEntityByClass(remarkIdentifier)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    r.remarkStorage,
		classRepository:  r.classRepository,
		detailRepository: r.detailRepository,
		factRepository:   r.factRepository,
	}, nil
}
