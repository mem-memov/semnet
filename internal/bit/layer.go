package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/class"
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
			// zero

			zeroClassIdentifier, err := classEntity.CreateBit()
			if err != nil {
				return Entity{}, Entity{}, err
			}

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

			oneClassIdentifier, err := classEntity.CreateBit()
			if err != nil {
				return Entity{}, Entity{}, err
			}

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
			// zero

			_, zeroCharacterIdentifier, err := l.storage.GetReference(classIdentifiers[0])
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.zeroEntity = l.entities.create(false, classIdentifiers[0], zeroCharacterIdentifier)

			// one

			_, oneCharacterIdentifier, err := l.storage.GetReference(classIdentifiers[1])
			if err != nil {
				return Entity{}, Entity{}, err
			}

			l.oneEntity = l.entities.create(true, classIdentifiers[1], oneCharacterIdentifier)
		default:
			return Entity{}, Entity{}, fmt.Errorf("wrong number of bit nodes in class %d", len(classIdentifiers))
		}

		l.isInitialized = true
	}

	return l.zeroEntity, l.oneEntity, nil
}
