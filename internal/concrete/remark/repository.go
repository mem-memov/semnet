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

func (r *Repository) CreateFirstUserRemark(object string, property string) (Aggregate, error) {

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return Aggregate{}, err
	}

	remark, err := r.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return Aggregate{}, err
	}

	// detail

	detail, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return Aggregate{}, err
	}

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return Aggregate{}, err
	}

	// fact

	fact, err := r.factRepository.CreateFirstStoryFact()
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToFact(fact.GetRemark())
	if err != nil {
		return Aggregate{}, err
	}

	err = fact.PointToRemark(remark.GetFact())
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    r.remarkStorage,
		classRepository:  r.classRepository,
		detailRepository: r.detailRepository,
		factRepository:   r.factRepository,
	}, nil
}

func (r *Repository) GetRemark(remarkIdentifier uint) (Aggregate, error) {

	return Aggregate{}, nil
}
