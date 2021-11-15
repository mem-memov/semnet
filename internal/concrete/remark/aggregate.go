package remark

import (
	"fmt"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Aggregate struct {
	remark           abstractRemark.Entity
	remarkStorage    abstractRemark.Storage
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

	nextRemarkIdentifier, err := a.remark.GetNextRemark()
	if err != nil {
		return Aggregate{}, err
	}

	nextRemark, err := a.remarkStorage.ReadEntityByPosition(nextRemarkIdentifier)

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	isRemark, err := class.IsRemark(nextRemark.GetClass())
	if err != nil {
		return Aggregate{}, err
	}

	if !isRemark {
		return Aggregate{}, fmt.Errorf("remark has wrong class: %d", nextRemark.GetClass())
	}

	return Aggregate{
		remark:           nextRemark,
		remarkStorage:    a.remarkStorage,
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

	nextFact, err := fact.ToNextFact()
	if err != nil {
		return Aggregate{}, err
	}

	remarkFact, err := nextFact.GetFirstRemark()
	if err != nil {
		return Aggregate{}, err
	}

	remark, err := a.remarkStorage.ReadEntityByFact(remarkFact)
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) HasNextStory() (bool, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return false, err
	}

	return fact.HasNextStory()
}

func (a Aggregate) GetNextStory() (Aggregate, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return Aggregate{}, err
	}

	nextFact, err := fact.ToNextStory()
	if err != nil {
		return Aggregate{}, err
	}

	nextIdentifier, err := nextFact.GetFirstRemark()
	if err != nil {
		return Aggregate{}, err
	}

	nextRemark, err := a.remarkStorage.ReadEntityByFact(nextIdentifier)
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           nextRemark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) GetObjectAndProperty() (string, string, error) {

	detailIdentifier, err := a.remark.GetSourceDetail()
	if err != nil {
		return "", "", err
	}

	detail, err := a.detailRepository.Fetch(detailIdentifier)
	if err != nil {
		return "", "", err
	}

	return detail.GetObjectAndProperty()
}

func (a Aggregate) AddRemarkToFact(property string) (Aggregate, error) {

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return Aggregate{}, err
	}

	remark, err := a.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return Aggregate{}, err
	}

	// detail

	detailIdentifier, err := a.remark.GetSourceDetail()
	if err != nil {
		return Aggregate{}, err
	}

	aggregateDetail, err := a.detailRepository.Fetch(detailIdentifier)
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

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return Aggregate{}, err
	}

	// fact

	factIdentifier, err := a.remark.GetTargetFact()
	if err != nil {
		return Aggregate{}, err
	}

	fact, err := a.factRepository.FetchByRemark(factIdentifier)
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToFact(fact.GetRemark())
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) AddFactToStory(object string, property string) (Aggregate, error) {

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return Aggregate{}, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return Aggregate{}, err
	}

	remark, err := a.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return Aggregate{}, err
	}

	// detail

	detail, err := a.detailRepository.Provide(object, property)
	if err != nil {
		return Aggregate{}, err
	}

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return Aggregate{}, err
	}

	// fact

	nextFact, err := a.factRepository.CreateNextFact(a.remark.GetFact())
	if err != nil {
		return Aggregate{}, err
	}

	err = remark.PointToFact(nextFact.GetRemark())
	if err != nil {
		return Aggregate{}, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}
