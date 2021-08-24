package node

type Phrases struct {
	storage storage
}

func NewPhrases(storage storage) *Phrases {
	return &Phrases{
		storage: storage,
	}
}

func (p *Phrases) Create(identifier uint) Phrase {
	return newPhrase(identifier, p.storage)
}
