package detail

type EntityMock struct {
	HasPhraseValue_       func(phraseValue string) (bool, error)
	AddRemark_            func(uint) error
	GetObjectAndProperty_ func() (string, string, error)
}

func (e EntityMock) HasPhraseValue(phraseValue string) (bool, error) {
	return e.HasPhraseValue_(phraseValue)
}

func (e EntityMock) AddRemark(remarkIdentifier uint) error {
	return e.AddRemark_(remarkIdentifier)
}

func (e EntityMock) GetObjectAndProperty() (string, string, error) {
	return e.GetObjectAndProperty_()
}
