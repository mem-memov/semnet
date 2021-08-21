package node

type Character struct {
	identifier uint
	storage    storage
}

func newCharacter(identifier uint, storage storage) Character {
	return Character{
		identifier: identifier,
		storage:    storage,
	}
}
