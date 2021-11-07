package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Storage struct {
	storage abstract.Storage
}

var _ abstractPhrase.Storage = &Storage{}

func NewStorage(storage abstract.Storage) *Storage {
	return &Storage{
		storage: storage,
	}
}

func (s *Storage) CreateEntity(
	classEntity abstractClass.Entity,
	wordEntity abstractWord.Aggregate,
) (abstractPhrase.Entity, error) {

	word, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	class, err := classEntity.CreatePhrase()
	if err != nil {
		return nil, err
	}

	phrase, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	detail, err := s.storage.Create()
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(class, word)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(word, phrase)
	if err != nil {
		return nil, err
	}

	err = s.storage.SetReference(phrase, detail)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByClass(class uint) (abstractPhrase.Entity, error) {

	_, word, err := s.storage.GetReference(class)
	if err != nil {
		return nil, err
	}

	_, phrase, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	_, detail, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByWord(word uint) (abstractPhrase.Entity, error) {

	class, phrase, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	_, detail, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByPhrase(phrase uint) (abstractPhrase.Entity, error) {

	word, detail, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: s.storage,
	}, nil
}

func (s *Storage) ReadEntityByDetail(detail uint) (abstractPhrase.Entity, error) {

	phrase, _, err := s.storage.GetReference(detail)
	if err != nil {
		return nil, err
	}

	word, _, err := s.storage.GetReference(phrase)
	if err != nil {
		return nil, err
	}

	class, _, err := s.storage.GetReference(word)
	if err != nil {
		return nil, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: s.storage,
	}, nil
}
