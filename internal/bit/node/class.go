package node

import "github.com/mem-memov/semnet/internal/class"

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
