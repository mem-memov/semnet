package fact

type Entity interface {
	Mark(sourceIdentifier uint) error
}
