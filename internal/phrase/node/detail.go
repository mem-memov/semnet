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
