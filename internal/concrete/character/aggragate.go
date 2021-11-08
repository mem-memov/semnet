package character

import (
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
)

type Aggregate struct {
	character         abstractCharacter.Entity
	characterElements abstractCharacter.Elements
	paths             *paths
}

var _ abstractCharacter.Aggregate = Aggregate{}

func (a Aggregate) HasTargetWord() (bool, error) {

	return a.character.HasTargetWord()
}

func (a Aggregate) PointToWord(word uint) error {

	return a.character.PointToWord(word)
}

func (a Aggregate) GetWord() uint {

	return a.character.GetWord()
}

func (a Aggregate) GetTargetWord() (uint, error) {

	return a.character.GetTargetWord()
}

func (a Aggregate) Extract() (rune, error) {

	characterElement := a.characterElements.CreateLastElement(a.character)

	bitValue, err := characterElement.ExtractBitValue()
	if err != nil {
		return 0, err
	}

	path := a.paths.create(bitValue)

	for {
		hasPreviousElement, err := characterElement.HasPreviousElement()
		if err != nil {
			return 0, err
		}

		if !hasPreviousElement {
			break
		}

		characterElement, err = characterElement.GetPreviousElement()

		bitValue, err := characterElement.ExtractBitValue()
		if err != nil {
			return 0, err
		}

		path = append(path, bitValue)
	}

	var integer int32

	for _, bitValue := range path.reverse() {
		integer <<= 1
		if bitValue {
			integer += 1
		}
	}

	return integer, nil
}
