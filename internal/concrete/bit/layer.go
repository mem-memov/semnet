package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type layer struct {
	storage         storage
	entities        *entities
	classRepository *class.Repository
	isInitialized   bool
	zeroEntity      Entity
	oneEntity       Entity
}

func newLayer(storage storage, entities *entities, classRepository *class.Repository) *layer {
	return &layer{
		storage:         storage,
		entities:        entities,
		classRepository: classRepository,
		isInitialized:   false,
	}
}

func (l *layer) initialize() (Entity, Entity, error) {

	if !l.isInitialized {
		classEntity, err := l.classRepository.ProvideEntity()
		if err != nil {
			return Entity{}, Entity{}, err
		}

		classIdentifiers, err := classEntity.GetAllBits()
		if err != nil {
			return Entity{}, Entity{}, err
		}

		switch len(classIdentifiers) {
		case 0:
			var classIdentifiers [2]uint
			var zeroClassIdentifier uint
			var oneClassIdentifier uint

			classIdentifiers[0], err = classEntity.CreateBit()
			if err != nil {
				return Entity{}, Entity{}, err
			}

			classIdentifiers[1], err = classEntity.CreateBit()
			if err != nil {
				return Entity{}, Entity{}, err
			}

			if classIdentifiers[0] < classIdentifiers[1] {
				zeroClassIdentifier = classIdentifiers[0]
				oneClassIdentifier = classIdentifiers[1]
			} else {
				zeroClassIdentifier = classIdentifiers[1]
				oneClassIdentifier = classIdentifiers[0]
			}

			// zero

			zeroCharacterIdentifier, err := l.storage.Create()
			if err != nil {
				return Entity{}, Entity{}, err
			}

			err = l.storage.SetReference(zeroClassIdentifier, zeroCharacterIdentifier)
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.zeroEntity = l.entities.create(false, zeroClassIdentifier, zeroCharacterIdentifier)

			// one

			oneCharacterIdentifier, err := l.storage.Create()
			if err != nil {
				return Entity{}, Entity{}, err
			}

			err = l.storage.SetReference(oneClassIdentifier, oneCharacterIdentifier)
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.oneEntity = l.entities.create(true, oneClassIdentifier, oneCharacterIdentifier)

		case 2:
			var zeroClassIdentifier uint
			var oneClassIdentifier uint

			if classIdentifiers[0] < classIdentifiers[1] {
				zeroClassIdentifier = classIdentifiers[0]
				oneClassIdentifier = classIdentifiers[1]
			} else {
				zeroClassIdentifier = classIdentifiers[1]
				oneClassIdentifier = classIdentifiers[0]
			}

			// zero

			_, zeroCharacterIdentifier, err := l.storage.GetReference(zeroClassIdentifier)
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.zeroEntity = l.entities.create(false, zeroClassIdentifier, zeroCharacterIdentifier)

			// one

			_, oneCharacterIdentifier, err := l.storage.GetReference(oneClassIdentifier)
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.oneEntity = l.entities.create(true, oneClassIdentifier, oneCharacterIdentifier)
		default:
			return Entity{}, Entity{}, fmt.Errorf("wrong number of bit nodes in class %d", len(classIdentifiers))
		}

		l.isInitialized = true
	}

	return l.zeroEntity, l.oneEntity, nil
}
