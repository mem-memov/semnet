package fact

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity struct {
	class    uint
	remark   uint
	position uint
	story    uint
	storage  abstract.Storage
}

var _ abstractFact.Entity = Entity{}

func createEntity(storage abstract.Storage) (Entity, error) {

	class, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	remark, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	position, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	story, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(class, remark)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(remark, position)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(position, story)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, remark, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	_, story, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  storage,
	}, nil
}

func readEntityByRemark(storage abstract.Storage, remark uint) (Entity, error) {

	class, position, err := storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	_, story, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  storage,
	}, nil
}

func readEntityByPosition(storage abstract.Storage, position uint) (Entity, error) {

	remark, story, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  storage,
	}, nil
}

func readEntityByStory(storage abstract.Storage, story uint) (Entity, error) {

	position, _, err := storage.GetReference(story)
	if err != nil {
		return Entity{}, err
	}

	remark, _, err := storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  storage,
	}, nil
}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetRemark() uint {

	return e.remark
}

func (e Entity) GetPosition() uint {

	return e.position
}

func (e Entity) GetStory() uint {

	return e.story
}

func (e Entity) PointToClass(class abstractClass.Entity) error {

	return e.storage.Connect(e.class, class.GetFact())
}

func (e Entity) PointToStory(story abstractStory.Entity) error {

	return e.storage.Connect(e.story, story.GetFact())
}

func (e Entity) PointToRemark(remark uint) error {

	return e.storage.Connect(e.remark, remark)
}

func (e Entity) HasNextFact() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetNextFact() (abstractFact.Entity, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return nil, err
	}

	if len(targets) != 1 {
		return nil, fmt.Errorf("fact has wrong number of next facts")
	}

	nextFact, err := readEntityByPosition(e.storage, targets[0])

	return nextFact, nil
}

func (e Entity) GetFirstRemark() (uint, error) {

	targets, err := e.storage.ReadTargets(e.remark)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("fact has wrong number of first remarks")
	}

	return targets[0], nil
}

func (e Entity) GetTargetStory() (uint, error) {

	targets, err := e.storage.ReadTargets(e.story)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("fact has wrong number of stories")
	}

	return targets[0], nil
}

func (e Entity) ToNextStory(nextFact uint) (abstractFact.Entity, error) {

	return readEntityByStory(e.storage, nextFact)
}
