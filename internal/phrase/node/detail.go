package node

type Detail struct {
	identifier uint
	storage    storage
}

func newDetail(identifier uint, storage storage) Detail {
	return Detail{
		identifier: identifier,
		storage:    storage,
	}
}