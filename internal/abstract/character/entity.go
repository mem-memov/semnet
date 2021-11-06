package character

type Entity interface {
	Mark(sourceIdentifier uint) error
	IsBeginningOfWords() (bool, error)
	ProvideSingleTarget() (uint, error)
}
