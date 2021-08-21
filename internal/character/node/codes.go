package node

import "github.com/mem-memov/semnet/internal/code"

type Codes struct {
	storage storage
	codeRepository *code.Repository
}

func NewCodes(storage storage, codeRepository *code.Repository) *Codes {
	return &Codes{
		storage: storage,
		codeRepository: codeRepository,
	}
}

func (c *Codes) Create(identifier uint) Code {
	return newCode(identifier, c.storage, c.codeRepository)
}
