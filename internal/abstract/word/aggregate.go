package word

type Aggregate interface {
	Extract() (string, error)
	IsBeginningOfPhrases() (bool, error)
	ProvideSingleTarget() (uint, error)
	Mark(phrase uint) error
}
