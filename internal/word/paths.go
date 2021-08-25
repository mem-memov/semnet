package word

import (
	"fmt"
)

type paths struct{}

func newPaths() *paths {
	return &paths{}
}

func (p *paths) create(start rune) path {

	newPath := make([]rune, 1)
	newPath[0] = start

	return newPath
}

func (p *paths) collect(word string) (path, error) {

	path := []rune(word)

	if len(path) < 1 {
		return path, fmt.Errorf("no character in word entity: %s", word)
	}

	return path, nil
}
