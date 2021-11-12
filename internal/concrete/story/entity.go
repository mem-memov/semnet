package story

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity struct {
	class    uint
	fact     uint
	position uint
	user     uint
	storage  abstract.Storage
}

var _ abstractStory.Entity = Entity{}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetFact() uint {

	return e.fact
}

func (e Entity) GetPosition() uint {

	return e.position
}

func (e Entity) GetUser() uint {

	return e.user
}

func (e Entity) PointToFact(fact uint) error {

	return e.storage.Connect(e.fact, fact)
}

func (e Entity) HasNextStory() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetTargetStory() (uint, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("story has wrong number of next stories")
	}

	return targets[0], nil
}

func (e Entity) GetTargetFact() (uint, error) {

	targets, err := e.storage.ReadTargets(e.fact)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("story has wrong number of target facts")
	}

	return targets[0], nil
}
