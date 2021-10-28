package story

type Entity interface {
	GetClass() uint
	GetFact() uint
	GetUser() uint
	PointToFact(fact uint) error
}
