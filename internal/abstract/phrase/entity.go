package phrase

type Entity interface {
	GetClass() uint
	GetWord() uint
	GetPhrase() uint
	GetDetail() uint

	PointToPhrase(phrase uint) error
	GetTargetPhrases() ([]uint, error)
	HasSourcePhrase() (bool, error)
	GetSourcePhrase() (uint, error)

	GetSourceWord() (uint, error)
	GetTargetWord() (uint, error)
	PointToWord(word uint) error

	GetSourceDetails() ([]uint, error)
	GetTargetDetails() ([]uint, error)
	AddSourceDetail(detail uint) error
	AddTargetDetail(detail uint) error

	Mark(sourceIdentifier uint) error
}
