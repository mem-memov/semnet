package fact

type Node interface {
	GetIdentifier() uint
	Mark(sourceIdentifier uint) error
}
