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

func createEntity(storage abstract.Storage) (Entity, error) {

	class, err := storage.Create()
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

func (e Entity) PointToClass(class abstractClass.Entity) error {

	return e.storage.Connect(e.fact, class.GetRemark())
}

func (e Entity) PointToPosition(remark abstractRemark.Entity) error {

	return e.storage.Connect(e.position, remark.GetPosition())
}

func (e Entity) PointToFact(fact abstractFact.Entity) error {

	return e.storage.Connect(e.fact, fact.GetRemark())
}

func (e Entity) FetchTargetFact(factRepository abstractFact.Repository) (abstractFact.Entity, error) {

	targetFacts, err := e.storage.ReadTargets(e.fact)
	if err != nil {
		return nil, err
	}

	if len(targetFacts) != 1 {
		return nil, fmt.Errorf("remark has wrong number of facts")
	}

	return factRepository.FetchByRemark(targetFacts[0])
}

func (e Entity) CreateNextStoryFact(factRepository abstractFact.Repository) (abstractFact.Entity, error) {

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

func (e Entity) GetNextRemark() (abstractRemark.Entity, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return nil, err
	}

	if len(targets) != 1 {
		return nil, fmt.Errorf("wrong number of next remarks")
	}

	nextRemark, err := readEntityByPosition(e.storage, targets[0])

	return nextRemark, nil
}
