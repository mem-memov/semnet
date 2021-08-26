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

func (d Detail) ProvideSingleTarget() (uint, error) {

	targets, err := d.storage.ReadTargets(d.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := d.storage.Create()
		if err != nil {
			return 0, err
		}

		err = d.storage.Connect(d.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

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
