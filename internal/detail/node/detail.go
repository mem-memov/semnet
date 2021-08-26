package node

import "fmt"

type Detail struct {
	identifier uint
	storage    storage
}

func newDetail(identifier uint, storage storage) Detail {
	return Detail{
		identifier: identifier,
		storage:    storage,
	}
}

func (d Detail) Identifier() uint {
	return d.identifier
}

func (d Detail) NewDetail(phrase Phrase) (Detail, error) {

	identifier, err := d.storage.Create()
	if err != nil {
		return Detail{}, err
	}

	err = d.storage.SetReference(phrase.Identifier(), identifier)
	if err != nil {
		return Detail{}, err
	}

	err = d.storage.Connect(d.identifier, identifier)
	if err != nil {
		return Detail{}, err
	}

	return newDetail(identifier, d.storage), nil
}

func (d Detail) ReadTargets() ([]Detail, error) {

	targetIdentifiers, err := d.storage.ReadTargets(d.identifier)
	if err != nil {
		return []Detail{}, err
	}

	if len(targetIdentifiers) > 2 {
		return []Detail{}, fmt.Errorf("too many targets in detail layer at detail %d", d.identifier)
	}

	details := make([]Detail, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		details[i] = newDetail(targetIdentifier, d.storage)
	}

	return details, nil
}

func (d Detail) ReadSources() ([]Detail, error) {

	sourceIdentifiers, err := d.storage.ReadSources(d.identifier)
	if err != nil {
		return []Detail{}, err
	}

	if len(sourceIdentifiers) > 1 {
		return []Detail{}, fmt.Errorf("too many sources in detail layer at detail %d", d.identifier)
	}

	details := make([]Detail, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		details[i] = newDetail(targetIdentifier, d.storage)
	}

	return details, nil
}

func (d Detail) GetPhraseAndRemark() (uint, uint, error) {

	phraseIdentifier, remarkIdentifier, err := d.storage.GetReference(d.identifier)
	if err != nil {
		return 0, 0, nil
	}

	return phraseIdentifier, remarkIdentifier, nil
}
