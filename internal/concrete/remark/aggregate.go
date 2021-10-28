package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Aggregate struct {
	remark           abstractRemark.Entity
	storage          abstract.Storage
	classRepository  abstractClass.Repository
	detailRepository abstractDetail.Repository
	factRepository   abstractFact.Repository
}

func (a Aggregate) GetFact() uint {
	return a.remark.GetFact()
}

func (a Aggregate) HasNextRemark() (bool, error) {

	return a.remark.HasNextRemark()
}

func (a Aggregate) GetNextRemark() (Aggregate, error) {

	remark, err := a.remark.GetNextRemark()
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) HasNextFact() (bool, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return false, err
	}

	return fact.HasNextFact()
}

func (a Aggregate) GetNextFact() (Aggregate, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return Aggregate{}, err
	}

	remarkFact, err := fact.GetFirstRemark()
	if err != nil {
		return Aggregate{}, err
	}

	remark, err := readEntityByFact(a.storage, remarkFact)
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) GetObjectAndProperty() (string, string, error) {

	detail, err := a.detailRepository.Fetch(a.remark.GetDetail())
	if err != nil {
		return "", "", err
	}

	return detail.GetObjectAndProperty()
}

func (a Aggregate) AddRemarkToFact(property string) (Aggregate, error) {

	remark, err := createEntity(a.storage)
	if err != nil {
		return Aggregate{}, err
	}

	// class

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToClass(class)
	if err != nil {
		return Aggregate{}, err
	}

	// detail

	aggregateDetail, err := a.detailRepository.Fetch(a.remark.GetDetail())
	if err != nil {
		return Aggregate{}, err
	}

	aggregateObject, _, err := aggregateDetail.GetObjectAndProperty()
	if err != nil {
		return Aggregate{}, err
	}

	detail, err := a.detailRepository.Provide(aggregateObject, property)
	if err != nil {
		return Aggregate{}, err
	}

	err = detail.PointToRemark(remark)
	if err != nil {
		return Aggregate{}, err
	}

	// fact

	fact, err := a.remark.FetchTargetFact(a.factRepository)

	err = remark.PointToFact(fact)
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) AddFactToStory(object string, property string) (Aggregate, error) {

	remark, err := createEntity(a.storage)
	if err != nil {
		return Aggregate{}, err
	}

	// class

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToClass(class)
	if err != nil {
		return Aggregate{}, err
	}

	// detail

	detail, err := a.detailRepository.Provide(object, property)
	if err != nil {
		return Aggregate{}, err
	}

	err = detail.PointToRemark(remark)
	if err != nil {
		return Aggregate{}, err
	}

	// position

	err = a.remark.PointToPosition(remark)
	if err != nil {
		return Aggregate{}, err
	}

	// fact

	fact, err := a.remark.CreateNextStoryFact(a.factRepository)
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToFact(fact)
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}
