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

	PointToWord(word uint) error
	HasTargetWord() (bool, error)
	GetTargetWord() (uint, error)
}
