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

var _ abstractRemark.Aggregate = Aggregate{}

func (a Aggregate) GetIdentifier() uint {
	return a.remark.GetClass()
}

func (a Aggregate) HasNextRemark() (bool, error) {

	return a.remark.HasNextRemark()
}

func (a Aggregate) GetNextRemark() (abstractRemark.Aggregate, error) {

	nextRemarkIdentifier, err := a.remark.GetNextRemark()
	if err != nil {
		return nil, err
	}

	nextRemark, err := a.remarkStorage.ReadEntityByPosition(nextRemarkIdentifier)

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	isRemark, err := class.IsRemark(nextRemark.GetClass())
	if err != nil {
		return nil, err
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

func (a Aggregate) GetNextFact() (abstractRemark.Aggregate, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	nextFact, err := fact.ToNextFact()
	if err != nil {
		return nil, err
	}

	remarkFact, err := nextFact.GetFirstRemark()
	if err != nil {
		return nil, err
	}

	remark, err := a.remarkStorage.ReadEntityByFact(remarkFact)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) GetFirstFact() (abstractRemark.Aggregate, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	firstFact, err := fact.ToFirstFact()
	if err != nil {
		return nil, err
	}

	remarkFact, err := firstFact.GetFirstRemark()
	if err != nil {
		return nil, err
	}

	remark, err := a.remarkStorage.ReadEntityByFact(remarkFact)
	if err != nil {
		return nil, err
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

func (a Aggregate) GetNextStory() (abstractRemark.Aggregate, error) {

	fact, err := a.factRepository.FetchByRemark(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	nextFact, err := fact.ToNextStory()
	if err != nil {
		return nil, err
	}

	nextIdentifier, err := nextFact.GetFirstRemark()
	if err != nil {
		return nil, err
	}

	nextRemark, err := a.remarkStorage.ReadEntityByFact(nextIdentifier)
	if err != nil {
		return nil, err
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

func (a Aggregate) AddRemarkToFact(property string) (abstractRemark.Aggregate, error) {

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return nil, err
	}

	remark, err := a.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	// detail

	detailIdentifier, err := a.remark.GetSourceDetail()
	if err != nil {
		return nil, err
	}

	aggregateDetail, err := a.detailRepository.Fetch(detailIdentifier)
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

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return nil, err
	}

	// fact

	factIdentifier, err := a.remark.GetTargetFact()
	if err != nil {
		return nil, err
	}

	fact, err := a.factRepository.FetchByRemark(factIdentifier)
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(fact.GetRemark())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) AddFactToStory(object string, property string) (abstractRemark.Aggregate, error) {

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return nil, err
	}

	remark, err := a.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	// detail

	detail, err := a.detailRepository.Provide(object, property)
	if err != nil {
		return nil, err
	}

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return nil, err
	}

	// fact

	nextFact, err := a.factRepository.CreateNextFact(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(nextFact.GetRemark())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}

func (a Aggregate) CreateChildStory(object string, property string) (abstractRemark.Aggregate, error) {

	class, err := a.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateRemark()
	if err != nil {
		return nil, err
	}

	remark, err := a.remarkStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	// detail

	detail, err := a.detailRepository.Provide(object, property)
	if err != nil {
		return nil, err
	}

	err = detail.PointToRemark(remark.GetDetail())
	if err != nil {
		return nil, err
	}

	// fact

	firstFact, err := a.factRepository.CreateChildStoryFact(a.remark.GetFact())
	if err != nil {
		return nil, err
	}

	err = remark.PointToFact(firstFact.GetRemark())
	if err != nil {
		return nil, err
	}

	err = firstFact.PointToRemark(remark.GetFact())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		remark:           remark,
		remarkStorage:    a.remarkStorage,
		classRepository:  a.classRepository,
		detailRepository: a.detailRepository,
		factRepository:   a.factRepository,
	}, nil
}
