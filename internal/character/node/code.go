package node

import "github.com/mem-memov/semnet/internal/code"

type Code struct {
	identifier     uint
	storage        storage
	codeRepository *code.Repository
}

func newCode(identifier uint, storage storage, codeRepository *code.Repository) Code {
	return Code{
		identifier:     identifier,
		storage:        storage,
		codeRepository: codeRepository,
	}
}
