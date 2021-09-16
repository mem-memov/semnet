package node

import "github.com/mem-memov/semnet/internal/class"

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

func (c *Classes) Create(identifier uint) Class {
	return newClass(identifier, c.storage, c.classRepository)
}
