package character

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractCharacter.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {
	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(
	classEntity abstractClass.Entity,
	bitEntity abstractBit.Entity,
) (abstractCharacter.Entity, error) {

	class, err := classEntity.CreateCharacter()
	if err != nil {
		return nil, err
	}

	bit, err := bitEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	character, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	word, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(class, bit)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(bit, character)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(character, word)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		bit:       bit,
		character: character,
		word:      word,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractCharacter.Entity, error) {

	_, bit, err := s.storage.GetReference(class)
	if err != nil {
		return nil, err
	}

	_, character, err := s.storage.GetReference(bit)
	if err != nil {
		return nil, err
	}

	_, word, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		bit:       bit,
		character: character,
		word:      word,
	}, nil
}

func (s *Storage) ReadEntityByBit(bit uint) (abstractCharacter.Entity, error) {

	class, character, err := s.storage.GetReference(bit)
	if err != nil {
		return nil, err
	}

	_, word, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		bit:       bit,
		character: character,
		word:      word,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByCharacter(character uint) (abstractCharacter.Entity, error) {

	bit, word, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(bit)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		bit:       bit,
		character: character,
		word:      word,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByWord(word uint) (abstractCharacter.Entity, error) {

	character, _, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	bit, _, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(bit)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		bit:       bit,
		character: character,
		word:      word,
		storage:   s.storage,
	}, nil
}
