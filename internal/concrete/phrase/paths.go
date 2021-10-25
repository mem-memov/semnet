package phrase

import (
	"fmt"
	"strings"
)

type paths struct{}

func newPaths() *paths {
	return &paths{}
}

func (p *paths) create(start string) path {

	newPath := make([]string, 1)
	newPath[0] = start

	return newPath
}

func (p *paths) collect(words string) (path, error) {

	path := strings.Split(words, " ")

	if len(path) < 1 {
		return path, fmt.Errorf("no words in phrase cluster: %s", words)
	}

	return path, nil
}
