package story

type Entity interface {
	GetClass() uint
	GetFact() uint
	GetPosition() uint
	GetUser() uint
	PointToFact(fact uint) error
	HasNextStory() (bool, error)
	GetNextStory() (Entity, error)
	GetTargetFact() (uint, error)
}
