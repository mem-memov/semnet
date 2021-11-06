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

func (w Word) IsBeginningOfWords() (bool, error) {

	target, err := w.ProvideSingleTarget()
	if err != nil {
		return false, err
	}

	backTargets, err := w.storage.ReadTargets(target)

	switch len(backTargets) {

	case 0:

		return false, nil

	case 1:

		if backTargets[0] != w.identifier {
			return false, fmt.Errorf("character not pointing to itself: %d", w.identifier)
		}

		return true, nil

	default:

		return false, fmt.Errorf("character not pointing to itself: %d", w.identifier)
	}
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
		return 0, fmt.Errorf("cluster %d has too many targets: %d", w.identifier, len(targets))
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

func (c Word) GetCharacter() (uint, error) {

	characterIdentifier, emptyReference, err := c.storage.GetReference(c.identifier)
	if err != nil {
		return 0, nil
	}

	if emptyReference != 0 {
		return 0, fmt.Errorf("next node reference is not empty in character layer at word %d", c.identifier)
	}

	return characterIdentifier, nil
}

func (w Word) String() string {
	return fmt.Sprintf("word %d", w.identifier)
}
