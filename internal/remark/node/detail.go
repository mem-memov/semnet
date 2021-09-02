package node

import "github.com/mem-memov/semnet/internal/detail"

type Detail struct {
	identifier       uint
	storage          storage
	detailRepository *detail.Repository
}

func newDetail(identifier uint, storage storage, detailRepository *detail.Repository) Detail {
	return Detail{
		identifier:       identifier,
		storage:          storage,
		detailRepository: detailRepository,
	}
}
