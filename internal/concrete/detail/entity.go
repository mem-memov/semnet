package detail

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Entity struct {
	class   uint
	phrase  uint
	remark  uint
	storage abstract.Storage
}

var _ abstractDetail.Entity = Entity{}

func createEntity(
	storage abstract.Storage,
	classEntity abstractClass.Entity,
	objectPhrase abstractPhrase.Entity,
	propertyPhrase abstractPhrase.Entity,
) (Entity, error) {

	var class, phrase, remark uint

	phrase, err := objectPhrase.ProvideDetailTarget(propertyPhrase)
	if err != nil {
		return Entity{}, err
	}

	class, remark, err = storage.GetReference(phrase)

	if class == 0 && remark == 0 {

		class, err = classEntity.CreateDetail()
		if err != nil {
			return Entity{}, err
		}

		remark, err = storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = storage.SetReference(class, phrase)
		if err != nil {
			return Entity{}, err
		}

		err = storage.SetReference(phrase, remark)
		if err != nil {
			return Entity{}, err
		}

	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, phrase, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, remark, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: storage,
	}, nil
}

func readEntityByPhrase(storage abstract.Storage, phrase uint) (Entity, error) {

	class, remark, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: storage,
	}, nil
}

func readEntityByRemark(storage abstract.Storage, remark uint) (Entity, error) {

	phrase, _, err := storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: storage,
	}, nil
}

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
