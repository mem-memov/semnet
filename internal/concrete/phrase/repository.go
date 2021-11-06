package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Repository struct {
	phraseStorage   abstractPhrase.Storage
	phraseElements  abstractPhrase.Elements
	paths           *paths
}

var _ abstractPhrase.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, wordRepository abstractWord.Repository) *Repository {

	phraseStorage := NewStorage(storage)
	phraseFactory := NewFactory(phraseStorage)

	return &Repository{
		phraseStorage:   phraseStorage,
		phraseElements:  NewElements(phraseStorage, phraseFactory, classRepository, wordRepository),
		paths:           newPaths(),
	}
}

func (r *Repository) Provide(words string) (abstractPhrase.Aggregate, error) {

	path, err := r.paths.collect(words)
	if err != nil {
		return nil, err
	}

	phraseElement, err := r.phraseElements.ProvideFirstElement(path[0])
	if err != nil {
		return nil, err
	}

	for _, wordValue := range path[1:] {

		phraseElement, err = phraseElement.ProvideNextElement(wordValue)
	}

	return Aggregate{
		phrase: phraseElement.GetEntity(),
		phraseElements: r.phraseElements,
		paths: r.paths,
	}, nil
}

func (r *Repository) Fetch(detail uint) (abstractPhrase.Aggregate, error) {

	phrase, err := r.phraseStorage.ReadEntityByDetail(detail)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		phrase: phrase,
		phraseElements: r.phraseElements,
		paths: r.paths,
	}, nil
}
