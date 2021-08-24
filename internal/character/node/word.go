package node

import "fmt"

type Word struct {
	identifier uint
	storage    storage
}

func newWord(identifier uint, storage storage) Word {
	return Word{
		identifier: identifier,
		storage:    storage,
	}
}

func (w Word) Identifier() uint {
	return w.identifier
}

func (w Word) Mark(sourceIdentifier uint) error {
	return w.storage.Connect(sourceIdentifier, w.identifier)
}

func (w Word) ProvideSingleTarget() (uint, error) {

	targets, err := w.storage.ReadTargets(w.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := w.storage.Create()
		if err != nil {
			return 0, err
		}

		err = w.storage.Connect(w.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", w.identifier, len(targets))
	}
}

func (w Word) NewWord(character Character) (Word, error) {

	identifier, err := w.storage.Create()
	if err != nil {
		return Word{}, nil
	}

	err = w.storage.SetReference(character.Identifier(), identifier)
	if err != nil {
		return Word{}, nil
	}

	return newWord(identifier, w.storage), nil
}

func (w Word) String() string {
	return fmt.Sprintf("word %d", w.identifier)
}
