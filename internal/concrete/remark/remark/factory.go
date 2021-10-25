package remark

import "github.com/mem-memov/semnet/internal/abstract"

type Remarks struct {
	storage abstract.Storage
}

func NewFactory(storage abstract.Storage) *Remarks {
	return &Remarks{
		storage: storage,
	}
}

func (r *Remarks) Create(identifier uint) Node {
	return newNode(identifier, r.storage)
}
