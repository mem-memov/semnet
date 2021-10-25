package node

type Characters struct {
	storage storage
}

func NewCharacters(storage storage) *Characters {
	return &Characters{
		storage: storage,
	}
}

func (c *Characters) Create(identifier uint) Character {
	return newCharacter(identifier, c.storage)
}
