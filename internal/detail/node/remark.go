package node

import "fmt"

type Remark struct {
	identifier uint
	storage    storage
}

func newRemark(identifier uint, storage storage) Remark {
	return Remark{
		identifier: identifier,
		storage:    storage,
	}
}

func (r Remark) NewRemark(phrase Phrase) (Remark, error) {

	identifier, err := r.storage.Create()
	if err != nil {
		return Remark{}, nil
	}

	err = r.storage.SetReference(phrase.Identifier(), identifier)
	if err != nil {
		return Remark{}, nil
	}

	return newRemark(identifier, r.storage), nil
}

func (r Remark) GetPhrase() (uint, error) {

	phraseIdentifier, emptyReference, err := r.storage.GetReference(r.identifier)
	if err != nil {
		return 0, nil
	}

	if emptyReference != 0 {
		return 0, fmt.Errorf("next node reference is not empty in detail layer at remark %d", r.identifier)
	}

	return phraseIdentifier, nil
}
