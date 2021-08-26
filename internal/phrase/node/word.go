package node

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/word"
)

type Word struct {
	identifier     uint
	storage        storage
	wordRepository *word.Repository
}

func newWord(identifier uint, storage storage, wordRepository *word.Repository) Word {
	return Word{
		identifier:     identifier,
		storage:        storage,
		wordRepository: wordRepository,
	}
}

func (w Word) Identifier() uint {
	return w.identifier
}

func (w Word) NewWord(wordValue string) (Word, error) {

	identifier, err := w.storage.Create()
	if err != nil {
		return Word{}, err
	}

	wordEntity, err := w.wordRepository.Provide(wordValue)
	if err != nil {
		return Word{}, err
	}

	err = wordEntity.Mark(identifier)
	if err != nil {
		return Word{}, err
	}

	return newWord(identifier, w.storage, w.wordRepository), nil
}

func (w Word) HasWordValue(value string) (bool, error) {

	targets, err := w.storage.ReadTargets(w.identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets %d in phrase layer at word %d", len(targets), w.identifier)
	}

	wordEntity, err := w.wordRepository.Fetch(targets[0])
	if err != nil {
		return false, err
	}

	wordValue, err := w.wordRepository.Extract(wordEntity)
	if err != nil {
		return false, err
	}

	return wordValue == value, nil
}

func (w Word) WordValue() (string, error) {

	targets, err := w.storage.ReadTargets(w.identifier)
	if err != nil {
		return "", err
	}

	if len(targets) != 1 {
		return "", fmt.Errorf("wrong number of targets %d in phrase layer at word %d", len(targets), w.identifier)
	}

	wordEntity, err := w.wordRepository.Fetch(targets[0])
	if err != nil {
		return "", err
	}

	wordValue, err := w.wordRepository.Extract(wordEntity)
	if err != nil {
		return "", err
	}

	return wordValue, nil
}
