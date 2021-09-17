package node

import (
	"fmt"
)

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

func (w Word) ReadTargets() ([]Word, error) {

	targetIdentifiers, err := w.storage.ReadTargets(w.identifier)
	if err != nil {
		return []Word{}, err
	}

	if len(targetIdentifiers) > 2 {
		return []Word{}, fmt.Errorf("too many targets in word layer at word %d", w.identifier)
	}

	words := make([]Word, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		words[i] = newWord(targetIdentifier, w.storage)
	}

	return words, nil
}

func (w Word) ReadSources() ([]Word, error) {

	sourceIdentifiers, err := w.storage.ReadSources(w.identifier)
	if err != nil {
		return []Word{}, err
	}

	if len(sourceIdentifiers) > 1 {
		return []Word{}, fmt.Errorf("too many sources in word layer at word %d", w.identifier)
	}

	words := make([]Word, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		words[i] = newWord(targetIdentifier, w.storage)
	}

	return words, nil
}

func (w Word) GetClassAndCharacterAndPhrase() (uint, uint, uint, error) {

	characterIdentifier, phraseIdentifier, err := w.storage.GetReference(w.identifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	classIdentifier, wordIdentifier, err := w.storage.GetReference(characterIdentifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	if wordIdentifier != w.identifier {
		return 0, 0, 0, fmt.Errorf("word entity invalid at word node %d", w.identifier)
	}

	return classIdentifier, characterIdentifier, phraseIdentifier, nil
}

func (w Word) NewWord(character Character) (Word, error) {

	identifier, err := w.storage.Create()
	if err != nil {
		return Word{}, err
	}

	err = w.storage.SetReference(character.Identifier(), identifier)
	if err != nil {
		return Word{}, err
	}

	err = w.storage.Connect(w.identifier, identifier)
	if err != nil {
		return Word{}, err
	}

	return newWord(identifier, w.storage), nil
}
