package character

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

type Repository struct {
	characterStorage  abstractCharacter.Storage
	characterElements abstractCharacter.Elements
	paths             *paths
}

var _ abstractCharacter.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, bitRepository abstractBit.Repository) *Repository {

	characterStorage := NewStorage(storage)
	characterFactory := NewFactory(characterStorage)

	return &Repository{
		characterStorage:  characterStorage,
		characterElements: NewElements(characterStorage, characterFactory, classRepository, bitRepository),
		paths:             newPaths(),
	}
}

func (r *Repository) Provide(integer rune) (abstractCharacter.Aggregate, error) {

	path, err := r.paths.collect(integer)
	if err != nil {
		return nil, err
	}

	characterElement, err := r.characterElements.ProvideFirstElement(path[0])
	if err != nil {
		return nil, err
	}

	for _, bitValue := range path[1:] {

		characterElement, err = characterElement.ProvideNextElement(bitValue)
	}

	return Aggregate{
		character:         characterElement.GetEntity(),
		characterElements: r.characterElements,
		paths:             r.paths,
	}, nil
}

func (r *Repository) Fetch(word uint) (abstractCharacter.Aggregate, error) {

	character, err := r.characterStorage.ReadEntityByWord(word)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		character:         character,
		characterElements: r.characterElements,
		paths:             r.paths,
	}, nil
}
