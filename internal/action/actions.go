package action

import (
	"github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/word"
	"strings"
)

type actions struct {
	storage semnet.storage
	words   *semnet.words
}

func newActions(storage semnet.storage, words *semnet.words) *actions {
	return &actions{
		storage: storage,
		words:   words,
	}
}

func (a *actions) create(name string) (Action, error) {
	wordNames := strings.Split(name, " ")

	words := make([]word.Word, len(wordNames))

	for i, wordName := range wordNames {
		word, err := a.words.create(wordName)
		if err != nil {
			return Action{}, err
		}
		words[i] = word
	}
}
