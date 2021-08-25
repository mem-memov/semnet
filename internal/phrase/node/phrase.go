package node

import "fmt"

type Phrase struct {
	identifier uint
	storage    storage
}

func newPhrase(identifier uint, storage storage) Phrase {
	return Phrase{
		identifier: identifier,
		storage:    storage,
	}
}

func (p Phrase) Identifier() uint {
	return p.identifier
}

func (p Phrase) NewPhrase(word Word) (Phrase, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return Phrase{}, err
	}

	err = p.storage.SetReference(word.Identifier(), identifier)
	if err != nil {
		return Phrase{}, err
	}

	err = p.storage.Connect(p.identifier, identifier)
	if err != nil {
		return Phrase{}, err
	}

	return newPhrase(identifier, p.storage), nil
}

func (p Phrase) ReadTargets() ([]Phrase, error) {

	targetIdentifiers, err := p.storage.ReadTargets(p.identifier)
	if err != nil {
		return []Phrase{}, err
	}

	if len(targetIdentifiers) > 2 {
		return []Phrase{}, fmt.Errorf("too many targets in phrase layer at phrase %d", p.identifier)
	}

	phrases := make([]Phrase, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		phrases[i] = newPhrase(targetIdentifier, p.storage)
	}

	return phrases, nil
}

func (p Phrase) ReadSources() ([]Phrase, error) {

	sourceIdentifiers, err := p.storage.ReadSources(p.identifier)
	if err != nil {
		return []Phrase{}, err
	}

	if len(sourceIdentifiers) > 1 {
		return []Phrase{}, fmt.Errorf("too many sources in phrase layer at phrase %d", p.identifier)
	}

	phrases := make([]Phrase, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		phrases[i] = newPhrase(targetIdentifier, p.storage)
	}

	return phrases, nil
}

func (p Phrase) GetWordAndDetail() (uint, uint, error) {

	wordIdentifier, detailIdentifier, err := p.storage.GetReference(p.identifier)
	if err != nil {
		return 0, 0, nil
	}

	return wordIdentifier, detailIdentifier, nil
}
