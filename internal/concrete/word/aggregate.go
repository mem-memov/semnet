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

		wordElement, err := wordElement.GetPreviousElement()

		characterValue, err := wordElement.ExtractCharacterValue()
		if err != nil {
			return "", err
		}

		path = append(path, characterValue)
	}

	return path.reverse().toString(), nil
}

func (a Aggregate) IsBeginningOfPhrases() (bool, error) {

	return a.word.IsBeginningOfPhrases()
}

func (a Aggregate) ProvideSingleTarget() (uint, error) {

	return a.word.ProvideSingleTarget()
}

func (a Aggregate) Mark(phrase uint) error {

	return a.word.Mark(phrase)
}
