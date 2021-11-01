package word

type Entity interface {
	PointToPhrase(phrase uint) error

	PhraseIdentifier() uint
	ProvideSingleTarget() (uint, error)
	Mark(sourceIdentifier uint) error
}
