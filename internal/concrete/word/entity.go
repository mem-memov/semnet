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

func (e Entity) PointToPhrase(phrase uint) error {

	return e.storage.Connect(e.phrase, phrase)
}

func (e Entity) Mark(sourceIdentifier uint) error {

	return e.storage.Connect(sourceIdentifier, e.phrase)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	targets, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := e.storage.Create()
		if err != nil {
			return 0, err
		}

		err = e.storage.Connect(e.phrase, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("word has wrong number of target phrases: %d at %d", len(targets), e.phrase)
	}
}

func (e Entity) IsBeginningOfPhrases() (bool, error) {

	target, err := e.ProvideSingleTarget()
	if err != nil {
		return false, err
	}

	backTargets, err := e.storage.ReadTargets(target)

	switch len(backTargets) {

	case 0:

		return false, nil

	case 1:

		if backTargets[0] != e.phrase {
			return false, fmt.Errorf("word not pointing to itself: %d", e.phrase)
		}

		return true, nil

	default:

		return false, fmt.Errorf("word not pointing to itself: %d", e.phrase)
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
		return 0, fmt.Errorf("too many sources in word tree: %d at %d", len(sourceWords), e.phrase)
	}

	return sourceWords[0], nil
}
