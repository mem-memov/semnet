package node

import (
	"fmt"
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
)

type Phrase struct {
	identifier uint
	storage    storage
}

var _ abstractNode.Phrase = &Phrase{}

func newPhrase(identifier uint, storage storage) abstractNode.Phrase {
	return Phrase{
		identifier: identifier,
		storage:    storage,
	}
}

func (p Phrase) Identifier() uint {
	return p.identifier
}

func (p Phrase) NewPhrase(word abstractNode.Word) (abstractNode.Phrase, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return nil, err
	}

	err = p.storage.SetReference(word.Identifier(), identifier)
	if err != nil {
		return nil, err
	}

	err = p.storage.Connect(p.identifier, identifier)
	if err != nil {
		return nil, err
	}

	return newPhrase(identifier, p.storage), nil
}

func (p Phrase) ReadTargets() ([]abstractNode.Phrase, error) {

	targetIdentifiers, err := p.storage.ReadTargets(p.identifier)
	if err != nil {
		return nil, err
	}

	if len(targetIdentifiers) > 2 {
		return nil, fmt.Errorf("too many targets in phrase layer at phrase %d", p.identifier)
	}

	phrases := make([]abstractNode.Phrase, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		phrases[i] = newPhrase(targetIdentifier, p.storage)
	}

	return phrases, nil
}

func (p Phrase) ReadSources() ([]abstractNode.Phrase, error) {

	sourceIdentifiers, err := p.storage.ReadSources(p.identifier)
	if err != nil {
		return nil, err
	}

	if len(sourceIdentifiers) > 1 {
		return nil, fmt.Errorf("too many sources in phrase layer at phrase %d", p.identifier)
	}

	phrases := make([]abstractNode.Phrase, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		phrases[i] = newPhrase(targetIdentifier, p.storage)
	}

	return phrases, nil
}

func (p Phrase) GetClassAndWordAndDetail() (uint, uint, uint, error) {

	wordIdentifier, detailIdentifier, err := p.storage.GetReference(p.identifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	classIdentifier, wordIdentifier, err := p.storage.GetReference(wordIdentifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	if wordIdentifier != p.identifier {
		return 0, 0, 0, fmt.Errorf("phrase cluster invalid at phrase node %d", p.identifier)
	}

	return classIdentifier, wordIdentifier, detailIdentifier, nil
}
