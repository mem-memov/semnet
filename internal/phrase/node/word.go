package node

import "github.com/mem-memov/semnet/internal/word"

type Word struct {
	identifier          uint
	storage             storage
	wordRepository *word.Repository
}

func newWord(identifier uint, storage storage, wordRepository *word.Repository) Word {
	return Word{
		identifier:          identifier,
		storage:             storage,
		wordRepository: wordRepository,
	}
}

func (w Word) Identifier() uint {
	return w.identifier
}