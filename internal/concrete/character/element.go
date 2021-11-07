package character

import (
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

type Element struct {
	character        abstractCharacter.Entity
	characterStorage abstractCharacter.Storage
	characterFactory abstractCharacter.Factory
	classRepository  abstractClass.Repository
	bitRepository    abstractBit.Repository
}

var _ abstractCharacter.Element = Element{}

func (e Element) GetEntity() abstractCharacter.Entity {

	return e.character
}

func (e Element) ProvideNextElement(bitValue bool) (abstractCharacter.Element, error) {

	targetCharacterIdentifiers, err := e.character.GetTargetCharacters()
	if err != nil {
		return nil, err
	}

	targetCharacters := make([]abstractCharacter.Entity, 0, len(targetCharacterIdentifiers))
	for _, targetCharacterIdentifier := range targetCharacterIdentifiers {

		targetCharacter, err := e.characterStorage.ReadEntityByCharacter(targetCharacterIdentifier)
		if err != nil {
			return nil, err
		}

		targetCharacters = append(targetCharacters, targetCharacter)
	}

	// use existing
	for _, targetCharacter := range targetCharacters {
		targetBitIdentifier, err := targetCharacter.GetTargetBit()
		if err != nil {
			return nil, err
		}

		targetBit, err := e.bitRepository.Fetch(targetBitIdentifier)
		if err != nil {
			return nil, err
		}

		targetValue := targetBit.Bit()

		if bitValue == targetValue {
			return Element{
				character:        targetCharacter,
				characterStorage: e.characterStorage,
				characterFactory: e.characterFactory,
				classRepository:  e.classRepository,
				bitRepository:    e.bitRepository,
			}, nil
		}
	}

	// create new
	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	bit, err := e.bitRepository.Provide(bitValue)
	if err != nil {
		return nil, err
	}

	newCharacter, err := e.characterFactory.ProvideFirstEntity(class, bit)
	if err != nil {
		return nil, err
	}

	err = e.character.PointToCharacter(newCharacter.GetCharacter())
	if err != nil {
		return nil, err
	}

	return Element{
		character:        newCharacter,
		characterStorage: e.characterStorage,
		characterFactory: e.characterFactory,
		classRepository:  e.classRepository,
		bitRepository:    e.bitRepository,
	}, nil
}

func (e Element) ExtractBitValue() (bool, error) {

	targetBitIdentifier, err := e.character.GetTargetBit()
	if err != nil {
		return false, err
	}

	bit, err := e.bitRepository.Fetch(targetBitIdentifier)
	if err != nil {
		return false, err
	}

	return bit.Bit(), nil
}

func (e Element) HasPreviousElement() (bool, error) {

	return e.character.HasSourceCharacter()
}

func (e Element) GetPreviousElement() (abstractCharacter.Element, error) {

	characterIdentifier, err := e.character.GetSourceCharacter()
	if err != nil {
		return nil, err
	}

	character, err := e.characterStorage.ReadEntityByCharacter(characterIdentifier)
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
