package fact

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
)

type Entity struct {
	class    uint
	remark   uint
	position uint
	story    uint
	storage  abstract.Storage
}

var _ abstractFact.Entity = Entity{}

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

func (e Entity) PointToStory(story uint) error {

	return e.storage.Connect(e.story, story)
}

func (e Entity) PointToRemark(remark uint) error {

	return e.storage.Connect(e.remark, remark)
}

func (e Entity) PointToPosition(position uint) error {

	return e.storage.Connect(e.position, position)
}

func (e Entity) HasTargetFact() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetTargetFact() (uint, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("fact has wrong number of target facts: %d at %d", len(targets), e.position)
	}

	return targets[0], nil
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
