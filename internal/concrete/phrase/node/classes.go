package node

import (
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Classes struct {
	storage         storage
	classRepository *class.Repository
}

func NewClasses(storage storage, classRepository *class.Repository) *Classes {
	return &Classes{
		storage:         storage,
		classRepository: classRepository,
	}
}

func (c *Classes) Create(identifier uint) abstractNode.Class {
	return newClass(identifier, c.storage, c.classRepository)
}

func (c *Classes) CreateNew() (abstractNode.Class, error) {
	classEntity, err := c.classRepository.ProvideEntity()
	if err != nil {
		return Class{}, err
	}

	identifier, err := classEntity.CreateCharacter()
	if err != nil {
		return Class{}, err
	}

	return newClass(identifier, c.storage, c.classRepository), nil
}
