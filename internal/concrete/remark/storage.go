package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractRemark.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {

	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(class uint) (abstractRemark.Entity, error) {

	detail, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	position, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	fact, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(class, detail)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(detail, position)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(position, fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractRemark.Entity, error) {

	_, detail, err := s.storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := s.storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	_, fact, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByDetail(detail uint) (abstractRemark.Entity, error) {

	class, position, err := s.storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	_, fact, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPosition(position uint) (abstractRemark.Entity, error) {

	detail, fact, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByFact(fact uint) (abstractRemark.Entity, error) {

	position, _, err := s.storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	detail, _, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		detail:   detail,
		position: position,
		fact:     fact,
		storage:  s.storage,
	}, nil
}
