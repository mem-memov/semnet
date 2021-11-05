package word

type Entity interface {
	PointToPhrase(phrase uint) error

	PhraseIdentifier() uint
	ProvideSingleTarget() (uint, error)
	HasSingleTargetOtherTargets() (bool, error)
	Mark(sourceIdentifier uint) error
}
