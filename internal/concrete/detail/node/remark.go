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

func (r Remark) GetClassAndPhrase() (uint, uint, error) {

	phraseIdentifier, emptyReference, err := r.storage.GetReference(r.identifier)
	if err != nil {
		return 0, 0, nil
	}

	if emptyReference != 0 {
		return 0, 0, fmt.Errorf("next node reference is not empty in detail layer at position %d", r.identifier)
	}

	classIdentifier, remarkIdentifier, err := r.storage.GetReference(phraseIdentifier)
	if err != nil {
		return 0, 0, nil
	}

	if remarkIdentifier != r.identifier {
		return 0, 0, fmt.Errorf("detail cluster invalid at position node %d", r.identifier)
	}

	return classIdentifier, phraseIdentifier, nil
}

func (r Remark) PointToRemark(targetIdentifier uint) error {

	return r.storage.Connect(r.identifier, targetIdentifier)
}