package word

import (
	"github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/character"
)

type words struct {
	storage    semnet.storage
	characters *semnet.characters
}

func newWords(storage semnet.storage, characters *semnet.characters) *words {
	return &words{
		storage:    storage,
		characters: characters,
	}
}

func (w *words) create(name string) (Word, error) {
	characters := make([]character.Entity, len([]rune(name)))

	for i, r := range name {
		characters[i] = w.characters.create(r)
	}
}
