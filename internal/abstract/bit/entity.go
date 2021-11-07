package bit

type Entity interface {
	Identifier() uint
	Is(bit bool) bool
	Bit() bool
	Mark(sourceIdentifier uint) error
	ProvideSingleTarget() (uint, error)
	IsBeginningOfCharacters() (bool, error)
}
