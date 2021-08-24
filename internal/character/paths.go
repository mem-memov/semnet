package character

import (
	"fmt"
)

type paths struct{}

func newPaths() *paths {
	return &paths{}
}

func (p *paths) create(start bool) path {

	newPath := make([]bool, 1)
	newPath[0] = start

	return newPath
}

func (p *paths) collect(integer rune) (path, error) {

	bitNames := fmt.Sprintf("%b", integer)

	if len(bitNames) < 1 {
		return path{}, fmt.Errorf("no bits in entity: %d", integer)
	}

	bitValues := path(make([]bool, len(bitNames)))

	for i, bitName := range bitNames {
		if bitName != '0' && bitName != '1' {
			return path{}, fmt.Errorf("invalid bit name: %c", bitName)
		}
		bitValues[i] = bitName == '1'
	}

	return bitValues, nil
}
