package node

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Phrase struct {
	identifier       uint
	storage          storage
	phraseRepository phrase.Repository
}

func newPhrase(identifier uint, storage storage, phraseRepository phrase.Repository) Phrase {
	return Phrase{
		identifier:       identifier,
		storage:          storage,
		phraseRepository: phraseRepository,
	}
}

func (p Phrase) Identifier() uint {
	return p.identifier
}

func (p Phrase) NewPhrase(wordValue string) (Phrase, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return Phrase{}, err
	}

	phraseEntity, err := p.phraseRepository.Provide(wordValue)
	if err != nil {
		return Phrase{}, err
	}

	err = phraseEntity.Mark(identifier)
	if err != nil {
		return Phrase{}, err
	}

	return newPhrase(identifier, p.storage, p.phraseRepository), nil
}

func (p Phrase) HasPhraseValue(value string) (bool, error) {

	targets, err := p.storage.ReadTargets(p.identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets %d in detail layer at phrase %d", len(targets), p.identifier)
	}

	phraseEntity, err := p.phraseRepository.Fetch(targets[0])
	if err != nil {
		return false, err
	}

	phraseValue, err := p.phraseRepository.Extract(phraseEntity)
	if err != nil {
		return false, err
	}

	return phraseValue == value, nil
}

func (p Phrase) PhraseValues() (string, string, error) {

	sources, err := p.storage.ReadSources(p.identifier)
	if err != nil {
		return "", "", err
	}

	if len(sources) != 1 {
		return "", "", fmt.Errorf("wrong number of sources %d in detail layer at phrase %d", len(sources), p.identifier)
	}

	objectPhraseEntity, err := p.phraseRepository.Fetch(sources[0])
	if err != nil {
		return "", "", err
	}

	objectPhraseValue, err := p.phraseRepository.Extract(objectPhraseEntity)
	if err != nil {
		return "", "", err
	}

	targets, err := p.storage.ReadTargets(p.identifier)
	if err != nil {
		return "", "", err
	}

	if len(targets) != 1 {
		return "", "", fmt.Errorf("wrong number of targets %d in detail layer at phrase %d", len(targets), p.identifier)
	}

	propertyPhraseEntity, err := p.phraseRepository.Fetch(targets[0])
	if err != nil {
		return "", "", err
	}

	propertyPhraseValue, err := p.phraseRepository.Extract(propertyPhraseEntity)
	if err != nil {
		return "", "", err
	}

	return objectPhraseValue, propertyPhraseValue, nil
}
