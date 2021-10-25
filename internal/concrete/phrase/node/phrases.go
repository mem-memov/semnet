package node

import abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"

type Phrases struct {
	storage storage
}

func NewPhrases(storage storage) *Phrases {
	return &Phrases{
		storage: storage,
	}
}

func (c *Phrases) Create(identifier uint) abstractNode.Phrase {
	return newPhrase(identifier, c.storage)
}
