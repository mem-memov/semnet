package semnet

import "strings"

type actions struct {
	words *words
}

func newActions(words *words) *actions {
	return &actions{
		words: words,
	}
}

func (a *actions) create(name string) (Action, error) {
	wordNames := strings.Split(name, " ")

	words := make([]Word, len(wordNames))

	for i, wordName := range wordNames {
		word, err := a.words.create(wordName)
		if err != nil {
			return Action{}, err
		}
		words[i] = word
	}
}
