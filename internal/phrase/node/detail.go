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

func (d Detail) NewDetail(phrase Phrase) (Detail, error) {

	identifier, err := d.storage.Create()
	if err != nil {
		return Detail{}, nil
	}

	err = d.storage.SetReference(phrase.Identifier(), identifier)
	if err != nil {
		return Detail{}, nil
	}

	return newDetail(identifier, d.storage), nil
}

func (d Detail) Identifier() uint {
	return d.identifier
}

func (d Detail) Mark(sourceIdentifier uint) error {
	return d.storage.Connect(sourceIdentifier, d.identifier)
}

func (d Detail) ProvideDetailTarget(another Detail) (uint, error) {

	targets, err := d.storage.ReadTargets(d.identifier)
	if err != nil {
		return 0, err
	}

	sources, err := d.storage.ReadSources(another.identifier)
	if err != nil {
		return 0, err
	}

	commonIdentifiers := make([]uint, 0, 1)

	// TODO: optimize (cut tails, use map, sort)
	for _, target := range targets {
		for _, source := range sources {
			if source == target {
				commonIdentifiers = append(commonIdentifiers, source)
			}
		}
	}

	switch len(commonIdentifiers) {

	case 0:
		commonIdentifier, err := d.storage.Create()
		if err != nil {
			return 0, err
		}

		err = d.storage.Connect(d.identifier, commonIdentifier)
		if err != nil {
			return 0, err
		}

		err = d.storage.Connect(commonIdentifier, another.identifier)
		if err != nil {
			return 0, err
		}

		return commonIdentifier, nil

	case 1:
		return commonIdentifiers[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", d.identifier, len(targets))
	}
}

func (d Detail) GetPhrase() (uint, error) {

	phraseIdentifier, emptyReference, err := d.storage.GetReference(d.identifier)
	if err != nil {
		return 0, nil
	}

	if emptyReference != 0 {
		return 0, fmt.Errorf("next node reference is not empty in phrase layer at detail %d", d.identifier)
	}

	return phraseIdentifier, nil
}
