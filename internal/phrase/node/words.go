package node

import "github.com/mem-memov/semnet/internal/word"

type Words struct {
	storage             storage
	wordRepository *word.Repository
}

func NewWords(storage storage, wordRepository *word.Repository) *Words {
	return &Words{
		storage:             storage,
		wordRepository: wordRepository,
	}
}

func (c *Words) Create(identifier uint) Word {
	return newWord(identifier, c.storage, c.wordRepository)
}