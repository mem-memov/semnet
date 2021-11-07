package node

import (
	"github.com/mem-memov/semnet/internal/abstract"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Classes struct {
	storage         abstract.Storage
	classRepository *class.Repository
}

func NewClasses(storage abstract.Storage, classRepository *class.Repository) *Classes {
	return &Classes{
		storage:         storage,
		classRepository: classRepository,
	}
}

func (c *Classes) Create(identifier uint) Class {
	return newClass(identifier, c.storage, c.classRepository)
}
