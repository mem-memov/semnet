package bit

type Entity interface {
	GetCharacter() uint
	Is(bit bool) bool
	Bit() bool
	MarkCharacter(sourceIdentifier uint) error
	ProvideSingleTarget() (uint, error)
	IsBeginningOfCharacters() (bool, error)
}
