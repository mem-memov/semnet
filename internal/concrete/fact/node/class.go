package node

import (
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Class struct {
	identifier      uint
	storage         storage
	classRepository *class.Repository
}

func newClass(identifier uint, storage storage, classRepository *class.Repository) Class {
	return Class{
		identifier:      identifier,
		storage:         storage,
		classRepository: classRepository,
	}
}

func (c Class) IsValid() (bool, error) {
	classEntity, err := c.classRepository.ProvideEntity()
	if err != nil {
		return false, err
	}

	return classEntity.IsCharacter(c.identifier)
}

func (c Class) NewClass() (Class, error) {
	classEntity, err := c.classRepository.ProvideEntity()
	if err != nil {
		return Class{}, err
	}

	newIdentifier, err := classEntity.CreateCharacter()
	if err != nil {
		return Class{}, err
	}

	return newClass(newIdentifier, c.storage, c.classRepository), nil
}
