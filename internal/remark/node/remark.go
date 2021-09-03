package node

type Remark struct {
	identifier uint
	storage    storage
}

func newRemark(identifier uint, storage storage) Remark {
	return Remark{
		identifier: identifier,
		storage:    storage,
	}
}
