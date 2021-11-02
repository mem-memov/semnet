package phrase

type Entity interface {
	GetClass() uint
	GetWord() uint
	GetPhrase() uint
	GetDetail() uint
	PointToPhrase(phrase uint) error
	GetTargetPhrases() ([]Entity, error)
	GetSourceWord() (uint, error)
	HasSourcePhrase() (bool, error)
	GetSourcePhrase() (Entity, error)
	GetSourceDetails() ([]uint, error)
	GetTargetDetails() ([]uint, error)
	AddSourceDetail(detail uint) error
	AddTargetDetail(detail uint) error
	Mark(sourceIdentifier uint) error
}
