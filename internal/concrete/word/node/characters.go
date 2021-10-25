package node

import (
	"github.com/mem-memov/semnet/internal/concrete/character"
)

type Characters struct {
	storage             storage
	characterRepository *character.Repository
}

func NewCharacters(storage storage, characterRepository *character.Repository) *Characters {
	return &Characters{
		storage:             storage,
		characterRepository: characterRepository,
	}
}

func (c *Characters) Create(identifier uint) Character {
	return newCharacter(identifier, c.storage, c.characterRepository)
}
