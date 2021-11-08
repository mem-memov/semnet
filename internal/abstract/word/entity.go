package word

type Entity interface {
	GetClass() uint
	GetCharacter() uint
	GetWord() uint
	GetPhrase() uint

	GetSourceCharacter() (uint, error)
	PointToCharacter(character uint) error
	GetTargetCharacter() (uint, error)

	HasSourceWord() (bool, error)
	GetSourceWord() (uint, error)
	PointToWord(word uint) error
	GetTargetWords() ([]uint, error)

	PointToPhrase(phrase uint) error
	GetTargetPhrase() (uint, error)
	HasTargetPhrase() (bool, error)
}
