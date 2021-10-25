package node

import "github.com/mem-memov/semnet/internal/abstract/phrase"

type Phrases struct {
	storage          storage
	phraseRepository phrase.Repository
}

func NewPhrases(storage storage, phraseRepository phrase.Repository) *Phrases {
	return &Phrases{
		storage:          storage,
		phraseRepository: phraseRepository,
	}
}

func (c *Phrases) Create(identifier uint) Phrase {
	return newPhrase(identifier, c.storage, c.phraseRepository)
}
