package node

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/character"
)

type Character struct {
	identifier          uint
	storage             storage
	characterRepository *character.Repository
}

func newCharacter(identifier uint, storage storage, characterRepository *character.Repository) Character {
	return Character{
		identifier:          identifier,
		storage:             storage,
		characterRepository: characterRepository,
	}
}

func (c Character) Identifier() uint {
	return c.identifier
}

func (c Character) NewCharacter(characterValue rune) (Character, error) {

	identifier, err := c.storage.Create()
	if err != nil {
		return Character{}, err
	}

	bitEntity, err := c.characterRepository.Provide(characterValue)
	if err != nil {
		return Character{}, err
	}

	err = bitEntity.Mark(identifier)
	if err != nil {
		return Character{}, err
	}

	return newCharacter(identifier, c.storage, c.characterRepository), nil
}

func (c Character) HasCharacterValue(value rune) (bool, error) {

	targets, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return false, err
	}

	if len(targets) != 1 {
		return false, fmt.Errorf("wrong number of targets %d in word layer at character %d", len(targets), c.identifier)
	}

	characterEntity, err := c.characterRepository.Fetch(targets[0])
	if err != nil {
		return false, err
	}

	characterValue, err := c.characterRepository.Extract(characterEntity)
	if err != nil {
		return false, err
	}

	return characterValue == value, nil
}

func (c Character) CharacterValue() (rune, error) {

	targets, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return 0, err
	}

	if len(targets) != 1 {
		return 0, fmt.Errorf("wrong number of targets %d in word layer at character %d", len(targets), c.identifier)
	}

	characterEntity, err := c.characterRepository.Fetch(targets[0])
	if err != nil {
		return 0, err
	}

	characterValue, err := c.characterRepository.Extract(characterEntity)
	if err != nil {
		return 0, err
	}

	return characterValue, nil
}
