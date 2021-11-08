package word

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractWord.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {
	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(class uint) (abstractWord.Entity, error) {

	character, err := s.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	word, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	phrase, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(class, character)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(character, word)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(word, phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		character: character,
		word:      word,
		phrase:    phrase,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractWord.Entity, error) {

	_, character, err := s.storage.GetReference(class)
	if err != nil {
		return nil, err
	}

	_, word, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	_, phrase, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		character: character,
		word:      word,
		phrase:    phrase,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByCharacter(character uint) (abstractWord.Entity, error) {

	class, word, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	_, phrase, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		character: character,
		word:      word,
		phrase:    phrase,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByWord(word uint) (abstractWord.Entity, error) {

	character, phrase, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		character: character,
		word:      word,
		phrase:    phrase,
		storage:   s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPhrase(phrase uint) (abstractWord.Entity, error) {

	word, _, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	character, _, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(character)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:     class,
		character: character,
		word:      word,
		phrase:    phrase,
		storage:   s.storage,
	}, nil
}
