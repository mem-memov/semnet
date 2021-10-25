package remark

type Node struct {
	identifier uint
	storage    storage
}

func newNode(identifier uint, storage storage) Node {
	return Node{
		identifier: identifier,
		storage:    storage,
	}
}
