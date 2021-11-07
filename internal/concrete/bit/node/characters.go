package node

import "github.com/mem-memov/semnet/internal/abstract"

type Characters struct {
	storage abstract.Storage
}

func NewCharacters(storage abstract.Storage) *Characters {
	return &Characters{
		storage: storage,
	}
}

func (c *Characters) Create(identifier uint) Character {
	return newCharacter(identifier, c.storage)
}
