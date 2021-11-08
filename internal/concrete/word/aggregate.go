package word

import (
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Aggregate struct {
	word         abstractWord.Entity
	wordElements abstractWord.Elements
	paths        *paths
}

var _ abstractWord.Aggregate = Aggregate{}

func (a Aggregate) Extract() (string, error) {

	wordElement := a.wordElements.CreateLastElement(a.word)

	characterValue, err := wordElement.ExtractCharacterValue()
	if err != nil {
		return "", err
	}

	path := a.paths.create(characterValue)

	for {
		hasPreviousElement, err := wordElement.HasPreviousElement()
		if err != nil {
			return "", err
		}

		if !hasPreviousElement {
			break
		}

		wordElement, err = wordElement.GetPreviousElement()

		characterValue, err := wordElement.ExtractCharacterValue()
		if err != nil {
			return "", err
		}

		path = append(path, characterValue)
	}

	return path.reverse().toString(), nil
}

func (a Aggregate) HasTargetPhrase() (bool, error) {

	return a.word.HasTargetPhrase()
}

func (a Aggregate) GetTargetPhrase() (uint, error) {

	return a.word.GetTargetPhrase()
}

func (a Aggregate) GetPhrase() uint {

	return a.word.GetPhrase()
}

func (a Aggregate) PointToPhrase(phrase uint) error {

	return a.word.PointToPhrase(phrase)
}
