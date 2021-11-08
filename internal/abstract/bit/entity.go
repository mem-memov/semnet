package bit

type Entity interface {
	GetCharacter() uint
	HasTargetCharacter() (bool, error)
	PointToCharacter(character uint) error
	GetTargetCharacter() (uint, error)
	Is(bit bool) bool
	Bit() bool
}
