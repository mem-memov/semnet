package node

type Fact struct {
	identifier uint
	storage    storage
}

func newFact(identifier uint, storage storage) Fact {
	return Fact{
		identifier: identifier,
		storage:    storage,
	}
}