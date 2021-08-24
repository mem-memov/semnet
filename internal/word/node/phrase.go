package node

type Phrase struct {
	identifier uint
	storage    storage
}

func newPhrase(identifier uint, storage storage) Phrase {
	return Phrase{
		identifier: identifier,
		storage: storage,
	}
}

func (p Phrase) NewPhrase(word Word) (Phrase, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return Phrase{}, nil
	}

	err = p.storage.SetReference(word.Identifier(), identifier)
	if err != nil {
		return Phrase{}, nil
	}

	return newPhrase(identifier, p.storage), nil
}
