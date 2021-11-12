package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractFact.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {

	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(class uint) (abstractFact.Entity, error) {

	remark, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	position, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	story, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(class, remark)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(remark, position)
	if err != nil {
		return Entity{}, err
	}

	err = s.storage.SetReference(position, story)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractFact.Entity, error) {

	_, remark, err := s.storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, position, err := s.storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	_, story, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByRemark(remark uint) (abstractFact.Entity, error) {

	class, position, err := s.storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	_, story, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPosition(position uint) (abstractFact.Entity, error) {

	remark, story, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  s.storage,
	}, nil
}

func (s *Storage) ReadEntityByStory(story uint) (abstractFact.Entity, error) {

	position, _, err := s.storage.GetReference(story)
	if err != nil {
		return Entity{}, err
	}

	remark, _, err := s.storage.GetReference(position)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := s.storage.GetReference(remark)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:    class,
		remark:   remark,
		position: position,
		story:    story,
		storage:  s.storage,
	}, nil
}
