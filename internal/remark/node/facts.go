package node

type Facts struct {
	storage storage
}

func NewFacts(storage storage) *Facts {
	return &Facts{
		storage: storage,
	}
}

func (f *Facts) Create(identifier uint) Fact {
	return newFact(identifier, f.storage)
}