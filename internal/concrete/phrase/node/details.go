package node

import abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"

type Details struct {
	storage storage
}

func NewDetails(storage storage) *Details {
	return &Details{
		storage: storage,
	}
}

func (d *Details) Create(identifier uint) abstractNode.Detail {
	return newDetail(identifier, d.storage)
}
