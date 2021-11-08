package word

type Aggregate interface {
	Extract() (string, error)
	HasTargetPhrase() (bool, error)
	GetTargetPhrase() (uint, error)
	GetPhrase() uint
	PointToPhrase(phrase uint) error
}
