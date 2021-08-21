package node

type Word struct {
	identifier uint
	storage    storage
}

func newWord(identifier uint, storage storage) Word {
	return Word{
		identifier: identifier,
		storage:    storage,
	}
}
