package character

type Entity interface {
	GetClass() uint
	GetBit() uint
	GetCharacter() uint
	GetWord() uint

	PointToBit(bit uint) error
	GetTargetBit() (uint, error)

	GetTargetCharacters() ([]uint, error)
	HasSourceCharacter() (bool, error)
	GetSourceCharacter() (uint, error)
	PointToCharacter(character uint) error

	MarkWord(sourceIdentifier uint) error
	IsBeginningOfWords() (bool, error)
	ProvideSingleTarget() (uint, error)
}
