package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Aggregate struct {
	detail           abstractDetail.Entity
	storage          abstract.Storage
	classRepository  abstractClass.Repository
	phraseRepository abstractPhrase.Repository
}

var _ abstractDetail.Aggregate = Aggregate{}

func (a Aggregate) ExtendObject() (abstractDetail.Aggregate, error) {
	return nil, nil
}

func (a Aggregate) ExtendProperty() (abstractDetail.Aggregate, error) {
	return nil, nil
}

func (a Aggregate) PointToRemark(remark uint) error {

	return a.detail.PointToRemark(remark)
}

func (a Aggregate) GetObjectAndProperty() (string, string, error) {

	objectIdentifier, propertyIdentifier, err := a.detail.GetObjectAndPropertyPhrases()
	if err != nil {
		return "", "", err
	}

	objectPhrase, err := a.phraseRepository.Fetch(objectIdentifier)
	if err != nil {
		return "", "", err
	}

	objectPhraseValue, err := a.phraseRepository.Extract(objectPhrase)
	if err != nil {
		return "", "", err
	}

	propertyPhrase, err := a.phraseRepository.Fetch(propertyIdentifier)
	if err != nil {
		return "", "", err
	}

	propertyPhraseValue, err := a.phraseRepository.Extract(propertyPhrase)
	if err != nil {
		return "", "", err
	}

	return objectPhraseValue, propertyPhraseValue, nil
}
