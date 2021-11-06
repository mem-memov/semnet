package word

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Repository struct {
	wordStorage  abstractWord.Storage
	wordFactory  abstractWord.Factory
	wordElements abstractWord.Elements
	paths        *paths
}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	characterRepository abstractCharacter.Repository,
) *Repository {

	wordStorage := NewStorage(storage)
	wordFactory := NewFactory(wordStorage)

	return &Repository{
		wordStorage:  wordStorage,
		wordFactory:  wordFactory,
		wordElements: NewElements(wordStorage, wordFactory, classRepository, characterRepository),
		paths:        newPaths(),
	}
}

func (r *Repository) Provide(word string) (abstractWord.Aggregate, error) {

	path, err := r.paths.collect(word)
	if err != nil {
		return nil, err
	}

	wordElement, err := r.wordElements.ProvideFirstElement(path[0])
	if err != nil {
		return nil, err
	}

	for _, characterValue := range path[1:] {

		wordElement, err = wordElement.ProvideNextElement(characterValue)
	}

	return Aggregate{
		word:         wordElement.GetEntity(),
		wordElements: r.wordElements,
		paths:        r.paths,
	}, nil
}

func (r *Repository) Fetch(phrase uint) (abstractWord.Aggregate, error) {

	word, err := r.wordStorage.ReadEntityByPhrase(phrase)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		word:         word,
		wordElements: r.wordElements,
		paths:        r.paths,
	}, nil
}
