package phrase

import abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"

type Aggregate struct {
	phrase         abstractPhrase.Entity
	phraseElements abstractPhrase.Elements
	paths          *paths
}

var _ abstractPhrase.Aggregate = Aggregate{}

func (a Aggregate) AddSourceDetail(detail uint) error {

	return a.phrase.AddSourceDetail(detail)
}

func (a Aggregate) AddTargetDetail(detail uint) error {

	return a.phrase.AddTargetDetail(detail)
}

func (a Aggregate) GetSourceDetails() ([]uint, error) {

	return a.phrase.GetSourceDetails()
}

func (a Aggregate) GetTargetDetails() ([]uint, error) {

	return a.phrase.GetTargetDetails()
}

func (a Aggregate) Extract() (string, error) {

	phraseElement := a.phraseElements.CreateLastElement(a.phrase)

	wordValue, err := phraseElement.ExtractWordValue()
	if err != nil {
		return "", err
	}

	path := a.paths.create(wordValue)

	for {
		hasPreviousElement, err := phraseElement.HasPreviousElement()
		if err != nil {
			return "", err
		}

		if !hasPreviousElement {
			break
		}

		phraseElement, err := phraseElement.GetPreviousElement()

		wordValue, err := phraseElement.ExtractWordValue()
		if err != nil {
			return "", err
		}

		path = append(path, wordValue)
	}

	return path.reverse().toString(), nil
}
