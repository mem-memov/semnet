package node

import (
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
	"github.com/mem-memov/semnet/internal/concrete/word"
)

type Words struct {
	storage        storage
	wordRepository *word.Repository
}

func NewWords(storage storage, wordRepository *word.Repository) *Words {
	return &Words{
		storage:        storage,
		wordRepository: wordRepository,
	}
}

func (c *Words) Create(identifier uint) abstractNode.Word {
	return newWord(identifier, c.storage, c.wordRepository)
}
