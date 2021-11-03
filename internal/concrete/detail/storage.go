package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractDetail.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {
	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(
	classEntity abstractClass.Entity,
	objectPhrase abstractPhrase.Entity,
	propertyPhrase abstractPhrase.Entity,
) (abstractDetail.Entity, error) {

	class, err := classEntity.CreateDetail()
	if err != nil {
		return nil, err
	}

	phrase, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	err = objectPhrase.AddTargetDetail(phrase)
	if err != nil {
		return nil, err
	}

	err = propertyPhrase.AddSourceDetail(phrase)
	if err != nil {
		return nil, err
	}

	remark, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(class, phrase)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(phrase, remark)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractDetail.Entity, error) {

	_, phrase, err := s.storage.GetReference(class)
	if err != nil {
		return nil, err
	}

	_, remark, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPhrase(phrase uint) (abstractDetail.Entity, error) {

	class, remark, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByRemark(remark uint) (abstractDetail.Entity, error) {

	phrase, _, err := s.storage.GetReference(remark)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		phrase:  phrase,
		remark:  remark,
		storage: s.storage,
	}, nil
}
