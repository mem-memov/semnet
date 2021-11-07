package node

import (
	"github.com/mem-memov/semnet/internal/abstract"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Class struct {
	identifier      uint
	storage         abstract.Storage
	classRepository *class.Repository
}

func newClass(identifier uint, storage abstract.Storage, classRepository *class.Repository) Class {
	return Class{
		identifier:      identifier,
		storage:         storage,
		classRepository: classRepository,
	}
}
