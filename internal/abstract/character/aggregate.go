package character

type Aggregate interface {
	HasTargetWord() (bool, error)
	PointToWord(word uint) error
	GetWord() uint
	GetTargetWord() (uint, error)
	Extract() (rune, error)
}
