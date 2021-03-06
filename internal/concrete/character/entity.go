package character

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
)

type Entity struct {
	class     uint
	bit       uint
	character uint
	word      uint
	storage   abstract.Storage
}

var _ abstractCharacter.Entity = Entity{}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetBit() uint {

	return e.bit
}

func (e Entity) GetCharacter() uint {

	return e.character
}

func (e Entity) GetWord() uint {

	return e.word
}

func (e Entity) PointToBit(bit uint) error {

	return e.storage.Connect(e.bit, bit)
}

func (e Entity) GetTargetBit() (uint, error) {

	targetBits, err := e.storage.ReadTargets(e.bit)
	if err != nil {
		return 0, err
	}

	if len(targetBits) != 1 {
		return 0, fmt.Errorf("character has wrong number of target bits: %d at %d", len(targetBits), e.bit)
	}

	return targetBits[0], nil
}

func (e Entity) GetTargetCharacters() ([]uint, error) {

	return e.storage.ReadTargets(e.character)
}

func (e Entity) HasSourceCharacter() (bool, error) {

	// TODO: read count from DB
	sourceCharacters, err := e.storage.ReadSources(e.character)
	if err != nil {
		return false, err
	}

	return len(sourceCharacters) != 0, nil
}

func (e Entity) GetSourceCharacter() (uint, error) {

	sourceCharacters, err := e.storage.ReadSources(e.character)
	if err != nil {
		return 0, err
	}

	if len(sourceCharacters) != 1 {
		return 0, fmt.Errorf("character has wrong number of source characters: %d at %d", len(sourceCharacters), e.character)
	}

	return sourceCharacters[0], nil
}

func (e Entity) PointToCharacter(character uint) error {

	return e.storage.Connect(e.character, character)
}

func (e Entity) PointToWord(word uint) error {

	return e.storage.Connect(e.word, word)
}

func (e Entity) MarkWord(sourceIdentifier uint) error {

	return e.storage.Connect(sourceIdentifier, e.word)
}

func (e Entity) HasTargetWord() (bool, error) {

	targets, err := e.storage.ReadTargets(e.character)
	if err != nil {
		return false, err
	}

	switch len(targets) {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("character has wrong number of target words: %d at %d", len(targets), e.character)
	}
}

func (e Entity) GetTargetWord() (uint, error) {

	targets, err := e.storage.ReadTargets(e.word)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("character has wrong number of target words: %d at %d", len(targets), e.word)
	}

	return targets[0], nil
}
