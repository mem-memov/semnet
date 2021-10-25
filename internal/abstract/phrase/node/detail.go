package node

type Detail interface {
	NewDetail(phrase Phrase) (Detail, error)
	Identifier() uint
	Mark(sourceIdentifier uint) error
	ProvideDetailTarget(another Detail) (uint, error)
	GetPhrase() (uint, error)
}