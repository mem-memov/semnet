package detail

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
)

type Entity struct {
	class   uint
	phrase  uint
	remark  uint
	storage abstract.Storage
}

var _ abstractDetail.Entity = Entity{}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetPhrase() uint {

	return e.phrase
}

func (e Entity) GetRemark() uint {

	return e.remark
}

func (e Entity) PointToRemark(remark uint) error {

	return e.storage.Connect(e.remark, remark)
}

func (e Entity) GetObjectAndPropertyPhrases() (uint, uint, error) {

	sources, err := e.storage.ReadSources(e.phrase)
	if err != nil {
		return 0, 0, err
	}

	if len(sources) != 1 {
		return 0, 0, fmt.Errorf("detail has wrong number of source phrases: %d at %d", len(sources), e.phrase)
	}

	targets, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return 0, 0, err
	}

	if len(targets) != 1 {
		return 0, 0, fmt.Errorf("detail has wrong number of target phrases: %d at %d", len(targets), e.phrase)
	}

	return sources[0], targets[0], nil
}
