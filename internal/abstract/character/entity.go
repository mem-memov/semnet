package character

type Entity interface {
	Mark(sourceIdentifier uint) error
	ProvideSingleTarget() (uint, error)
}
