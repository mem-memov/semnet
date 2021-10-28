package remark

import (
	api "github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type aggregate struct {
	remark abstractRemark.Entity
	storage abstract.Storage
	classRepository abstractClass.Repository
	detailRepository abstractDetail.Repository
	factRepository   abstractFact.Repository
}

func (a aggregate) GetFact() uint {
	return a.remark.GetFact()
}

func (a aggregate) HasNextRemark() (bool, error) {

	return a.remark.HasNextRemark()
}

func (a aggregate) GetNextRemark() (api.Remark, error) {

	remark, err := a.remark.GetNextRemark()
	if err != nil {
		return nil, err
	}

	return aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a aggregate) HasNextFact() (bool, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return false, err
	}

	return fact.HasNextFact()
}

func (a aggregate) GetNextFact() (api.Remark, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	remarkFact, err := fact.GetFirstRemark()
	if err != nil {
		return nil, err
	}

	remark, err := readEntityByFact(a.storage, remarkFact)
	if err != nil {
		return nil, err
	}

	return aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a aggregate) GetObjectAndProperty() (string, string, error) {

	detail, err := a.detailRepository.Fetch(a.remark.GetDetail())
	if err != nil {
		return "", "", err
	}

	return detail.GetObjectAndProperty()
}

func (a aggregate) AddRemarkToFact(property string) (api.Remark, error)  {

	remark, err := createEntity(a.storage)
	if err != nil {
		return nil, err
	}

	// class

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	err = remark.PointToClass(class)
	if err != nil {
		return nil, err
	}

	// detail

	aggregateDetail, err := a.detailRepository.Fetch(a.remark.GetDetail())
	if err != nil {
		return nil, err
	}

	aggregateObject, _, err := aggregateDetail.GetObjectAndProperty()
	if err != nil {
		return nil, err
	}

	detail, err := a.detailRepository.Provide(aggregateObject, property)
	if err != nil {
		return nil, err
	}

	err = detail.PointToRemark(remark)
	if err != nil {
		return nil, err
	}

	// fact

	fact, err := a.remark.FetchTargetFact(a.factRepository)

	err = remark.PointToFact(fact)
	if err != nil {
		return nil, err
	}

	return aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a aggregate) AddFactToStory(object string, property string) (api.Remark, error)  {

	remark, err := createEntity(a.storage)
	if err != nil {
		return nil, err
	}

	// class

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	err = remark.PointToClass(class)
	if err != nil {
		return nil, err
	}

	// detail

	detail, err := a.detailRepository.Provide(object, property)
	if err != nil {
		return nil, err
	}


	err = detail.PointToRemark(remark)
	if err != nil {
		return nil, err
	}

	// position

	err = a.remark.PointToPosition(remark)
	if err != nil {
		return nil, err
	}

	// fact

	fact, err := a.remark.CreateNextStoryFact(a.factRepository)
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(fact)
	if err != nil {
		return nil, err
	}

	return aggregate{
		remark:           remark,
		storage:          a.storage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}
