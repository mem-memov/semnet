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
	tree     uint
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

func (e Entity) GetTree() uint {

	return e.tree
}

func (e Entity) PointToFact(fact uint) error {

	return e.storage.Connect(e.fact, fact)
}

func (e Entity) PointToPosition(position uint) error {

	return e.storage.Connect(e.position, position)
}

func (e Entity) PointToTree(tree uint) error {

	return e.storage.Connect(e.tree, tree)
}

func (e Entity) HasTargetPosition() (bool, error) {

	targets, err := e.storage.ReadTargets(e.position)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetTargetPosition() (uint, error) {

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

func (e Entity) HasTargetTree() (bool, error) {

	targets, err := e.storage.ReadTargets(e.tree)
	if err != nil {
		return false, err
	}

	return len(targets) != 0, nil
}

func (e Entity) GetTargetTree() (uint, error) {

	targets, err := e.storage.ReadTargets(e.tree)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("story has wrong number of target trees")
	}

	return targets[0], nil
}

func (e Entity) HasSourceTree() (bool, error) {

	sources, err := e.storage.ReadSources(e.tree)
	if err != nil {
		return false, err
	}

	return len(sources) != 0, nil
}

func (e Entity) GetSourceTree() (uint, error) {

	sources, err := e.storage.ReadSources(e.tree)
	if err != nil {
		return 0, err
	}

	if len(sources) != 1 {
		return 0, fmt.Errorf("story has wrong number of source trees")
	}

	return sources[0], nil
}
