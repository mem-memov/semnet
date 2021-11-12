package story

type Entity interface {
	GetClass() uint
	GetFact() uint
	GetPosition() uint
	GetUser() uint

	PointToFact(fact uint) error
	HasNextStory() (bool, error)
	GetTargetStory() (uint, error)
	GetTargetFact() (uint, error)
}
