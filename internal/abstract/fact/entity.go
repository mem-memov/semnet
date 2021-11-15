package fact

type Entity interface {
	GetClass() uint
	GetRemark() uint
	GetPosition() uint
	GetStory() uint

	PointToStory(story uint) error
	PointToRemark(remark uint) error
	PointToPosition(position uint) error

	HasTargetFact() (bool, error)
	GetTargetFact() (uint, error)

	GetFirstRemark() (uint, error)
	GetTargetStory() (uint, error)
}
