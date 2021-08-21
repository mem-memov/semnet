package code

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
)

type paths struct {
	bitRepository *bit.Repository
}

func newPaths(bitRepository *bit.Repository) *paths {
	return &paths{
		bitRepository: bitRepository,
	}
}

func (p *paths) collect(integer int32) (path, error) {

	bitNames := fmt.Sprintf("%b", integer)

	if len(bitNames) < 1 {
		return path{}, fmt.Errorf("no bits in entity: %d", integer)
	}

	bitEntities := path(make([]bit.Entity, len(bitNames)))

	for i, bitName := range bitNames {
		if bitName != '0' && bitName != '1' {
			return path{}, fmt.Errorf("invalid bit name: %c", bitName)
		}

		bitEntity, err := p.bitRepository.Create(bitName == '1')
		if err != nil {
			return path{}, err
		}

		bitEntities[i] = bitEntity
	}

	return bitEntities, nil
}
