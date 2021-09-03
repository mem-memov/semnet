package node

type Topic struct {
	identifier uint
	storage    storage
}

func newTopic(identifier uint, storage storage) Topic {
	return Topic{
		identifier: identifier,
		storage:    storage,
	}
}
