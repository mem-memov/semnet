package detail

type Entity interface {
	HasPhraseValue(phraseValue string) (bool, error)
	AddRemark(remarkIdentifier uint) error
	GetObjectAndProperty() (string, string, error)
}
