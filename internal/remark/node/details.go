package node

import "github.com/mem-memov/semnet/internal/detail"

type Details struct {
	storage storage
	detailRepository *detail.Repository
}

func NewDetails(storage storage, detailRepository *detail.Repository) *Details {
	return &Details{
		storage: storage,
		detailRepository: detailRepository,
	}
}

func (d *Details) Create(identifier uint) Detail {
	return newDetail(identifier, d.storage, d.detailRepository)
}
