package character

import (
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

type Elements struct {
	characterStorage abstractCharacter.Storage
	characterFactory abstractCharacter.Factory
	classRepository  abstractClass.Repository
	bitRepository    abstractBit.Repository
}

var _ abstractCharacter.Elements = &Elements{}

func NewElements(
	characterStorage abstractCharacter.Storage,
	characterFactory abstractCharacter.Factory,
	classRepository abstractClass.Repository,
	bitRepository abstractBit.Repository,
) *Elements {

	return &Elements{
		characterStorage: characterStorage,
		characterFactory: characterFactory,
		classRepository:  classRepository,
		bitRepository:    bitRepository,
	}
}

func (e *Elements) ProvideFirstElement(bitValue bool) (abstractCharacter.Element, error) {

	firstBit, err := e.bitRepository.Provide(bitValue)
	if err != nil {
		return nil, err
	}

	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	character, err := e.characterFactory.ProvideFirstEntity(class, firstBit)

	err = firstBit.Mark(character.GetBit())
	if err != nil {
		return nil, err
	}

	return Element{
		character:        character,
		characterStorage: e.characterStorage,
		characterFactory: e.characterFactory,
		classRepository:  e.classRepository,
		bitRepository:    e.bitRepository,
	}, nil
}

func (e *Elements) CreateLastElement(character abstractCharacter.Entity) abstractCharacter.Element {

	return Element{
		character:        character,
		characterStorage: e.characterStorage,
		characterFactory: e.characterFactory,
		classRepository:  e.classRepository,
		bitRepository:    e.bitRepository,
	}
}
