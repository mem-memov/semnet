package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractStory.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {

	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(class uint) (abstractStory.Entity, error) {

	fact, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	position, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	tree, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(class, fact)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(fact, position)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(position, tree)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		fact:     fact,
		position: position,
		tree:     tree,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractStory.Entity, error) {

	_, fact, err := s.storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := s.storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	_, tree, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		fact:     fact,
		position: position,
		tree:     tree,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByFact(fact uint) (abstractStory.Entity, error) {

	class, position, err := s.storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	_, tree, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		fact:     fact,
		position: position,
		tree:     tree,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPosition(position uint) (abstractStory.Entity, error) {

	fact, tree, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		fact:     fact,
		position: position,
		tree:     tree,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByTree(tree uint) (abstractStory.Entity, error) {

	position, _, err := s.storage.GetReference(tree)
	if err != nil {
		return Entity{}, err
	}

	fact, _, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(fact)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		fact:     fact,
		position: position,
		tree:     tree,
		storage:  s.storage,
	}, nil
}
