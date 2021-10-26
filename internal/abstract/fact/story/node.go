package story

type Node interface {
	GetIdentifier() uint
	Mark(sourceIdentifier uint) error
}
