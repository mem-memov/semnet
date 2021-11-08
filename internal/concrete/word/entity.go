package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Entity struct {
	class     uint
	character uint
	word      uint
	phrase    uint
	storage   abstract.Storage
}

var _ abstractWord.Entity = Entity{}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetCharacter() uint {

	return e.character
}

func (e Entity) GetWord() uint {

	return e.word
}

func (e Entity) GetPhrase() uint {

	return e.phrase
}

func (e Entity) GetSourceCharacter() (uint, error) {

	sourceCharacters, err := e.storage.ReadSources(e.character)
	if err != nil {
		return 0, err
	}

	if len(sourceCharacters) != 1 {
		return 0, fmt.Errorf("word has wrong number of source characters: %d at %d", len(sourceCharacters), e.character)
	}

	return sourceCharacters[0], nil
}

func (e Entity) PointToCharacter(character uint) error {

	return e.storage.Connect(e.character, character)
}

func (e Entity) PointToPhrase(phrase uint) error {

	return e.storage.Connect(e.phrase, phrase)
}

func (e Entity) GetTargetPhrase() (uint, error) {

	targets, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("word has wrong number of target phrases: %d at %d", len(targets), e.phrase)
	}

	return targets[0], nil
}

func (e Entity) HasTargetPhrase() (bool, error) {

	targets, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return false, err
	}

	switch len(targets) {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("word has wrong number of target phrases: %d at %d", len(targets), e.character)
	}
}

func (e Entity) GetTargetWords() ([]uint, error) {

	return e.storage.ReadTargets(e.word)
}

func (e Entity) GetTargetCharacter() (uint, error) {

	targetCharacters, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return 0, err
	}

	if len(targetCharacters) != 1 {
		return 0, fmt.Errorf("word has wrong number of target charcters: %d at %d", len(targetCharacters), e.word)
	}

	return targetCharacters[0], nil
}

func (e Entity) HasSourceWord() (bool, error) {

	// TODO: read count from DB
	sourceWords, err := e.storage.ReadSources(e.word)
	if err != nil {
		return false, err
	}

	return len(sourceWords) != 0, nil
}

func (e Entity) GetSourceWord() (uint, error) {

	sourceWords, err := e.storage.ReadSources(e.word)
	if err != nil {
		return 0, err
	}

	if len(sourceWords) != 1 {
		return 0, fmt.Errorf("word has wrong number of source words: %d at %d", len(sourceWords), e.phrase)
	}

	return sourceWords[0], nil
}

func (e Entity) PointToWord(word uint) error {

	return e.storage.Connect(e.word, word)
}
