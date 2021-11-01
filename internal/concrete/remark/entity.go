package remark

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Entity struct {
	class    uint
	detail   uint
	position uint
	fact     uint
	storage  abstract.Storage
}

var _ abstractRemark.Entity = Entity{}

func createEntity(storage abstract.Storage, classEntity abstractClass.Entity) (Entity, error) {

	class, err := classEntity.CreateRemark()
	if err != nil {
		return Entity{}, err
	}

	detail, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	position, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	fact, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(class, detail)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(detail, position)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(position, fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, detail, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	_, fact, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  storage,
	}, nil
}

func readEntityByDetail(storage abstract.Storage, detail uint) (Entity, error) {

	class, position, err := storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	_, fact, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  storage,
	}, nil
}

func readEntityByPosition(storage abstract.Storage, position uint) (Entity, error) {

	detail, fact, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  storage,
	}, nil
}

func readEntityByFact(storage abstract.Storage, fact uint) (Entity, error) {

	position, _, err := storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	detail, _, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  storage,
	}, nil
}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetDetail() uint {

	return e.detail
}

func (e Entity) GetPosition() uint {

	return e.position
}

func (e Entity) GetFact() uint {

	return e.fact
}

func (e Entity) PointToPosition(remark abstractRemark.Entity) error {

	return e.storage.Connect(e.position, remark.GetPosition())
}

func (e Entity) PointToFact(fact abstractFact.Aggregate) error {

	return e.storage.Connect(e.fact, fact.GetRemark())
}

func (e Entity) GetTargetFact() (uint, error) {

	targetFacts, err := e.storage.ReadTargets(e.fact)
	if err != nil {
		return 0, err
	}

	if len(targetFacts) != 1 {
		return 0, fmt.Errorf("remark has wrong number of target facts")
	}

	return targetFacts[0], nil
}

func (e Entity) GetSourceDetail() (uint, error) {

	sourceDetails, err := e.storage.ReadSources(e.detail)
	if err != nil {
		return 0, err
	}

	if len(sourceDetails) != 1 {
		return 0, fmt.Errorf("remark has wrong number of source details")
	}

	return sourceDetails[0], nil
}

func (e Entity) CreateNextStoryFact(factRepository abstractFact.Repository) (abstractFact.Aggregate, error) {

	targetFacts, err := e.storage.ReadTargets(e.fact)
	if err != nil {
		return nil, err
	}

	if len(targetFacts) != 1 {
		return nil, fmt.Errorf("remark has wrong number of facts")
	}

	// TODO: create new fact of the same story
	return nil, nil
}

func (e Entity) HasNextRemark() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetNextRemark() (uint, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("wrong number of next remarks")
	}

	return targets[0], nil
}

func (e Entity) ToNextFact(fact uint) (abstractRemark.Entity, error) {

	return readEntityByFact(e.storage, fact)
}
