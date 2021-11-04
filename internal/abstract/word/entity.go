package word

type Entity interface {
	PointToPhrase(phrase uint) error

	PhraseIdentifier() uint
	ProvideSingleTarget() (uint, error)
	HasSingleTargetSources() (bool, error)
	Mark(sourceIdentifier uint) error
}
