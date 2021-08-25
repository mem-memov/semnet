package node

type Phrase struct {
	identifier uint
	storage    storage
}

func newPhrase(identifier uint, storage storage) Phrase {
	return Phrase{
		identifier: identifier,
		storage:    storage,
	}
}