package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/class"
)

type layer struct {
	storage         storage
	classRepository *class.Repository
	isInitialized   bool
	zeroIdentifier uint
	oneIdentifier uint
}

func newLayer(storage storage, classRepository *class.Repository) *layer {
	return &layer{
		storage:         storage,
		classRepository: classRepository,
		isInitialized:   false,
	}
}

func (l *layer) initialize() (uint, uint, error) {

	if !l.isInitialized {
		classEntity, err := l.classRepository.ProvideEntity()
		if err != nil {
			return 0, 0, err
		}

		bitIdentifiers, err := classEntity.GetAllBits()
		if err != nil {
			return 0, 0, err
		}

		switch len(bitIdentifiers) {
		case 0:
			l.zeroIdentifier, err = classEntity.CreateBit()
			if err != nil {
				return 0, 0, err
			}

			l.oneIdentifier, err = classEntity.CreateBit()
			if err != nil {
				return 0, 0, err
			}
		case 2:
			l.zeroIdentifier = bitIdentifiers[0]
			l.oneIdentifier = bitIdentifiers[1]
		default:
			return 0, 0, fmt.Errorf("wrong number of bit nodes in class %d", len(bitIdentifiers))
		}

		l.isInitialized = true
	}

	return l.zeroIdentifier, l.oneIdentifier, nil
}
