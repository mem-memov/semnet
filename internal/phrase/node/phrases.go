package node

type Phrases struct {
	storage storage
}

func NewPhrases(storage storage) *Phrases {
	return &Phrases{
		storage: storage,
	}
}

func (c *Phrases) Create(identifier uint) Phrase {
	return newPhrase(identifier, c.storage)
}
