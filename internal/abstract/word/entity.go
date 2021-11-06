package word

type Entity interface {
	GetClass() uint
	GetCharacter() uint
	GetWord() uint
	GetPhrase() uint

	GetSourceCharacter() (uint, error)
	GetTargetCharacter() (uint, error)

	PointToPhrase(phrase uint) error

	HasSourceWord() (bool, error)
	GetSourceWord() (uint, error)
	GetTargetWords() ([]uint, error)

	ProvideSingleTarget() (uint, error)
	IsBeginningOfPhrases() (bool, error)
	Mark(sourceIdentifier uint) error
}
