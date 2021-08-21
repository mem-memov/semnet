package node

type Words struct {
	storage storage
}

func NewWords(storage storage) *Words {
	return &Words{
		storage: storage,
	}
}

func (c *Words) Create(identifier uint) Word {
	return newWord(identifier, c.storage)
}
