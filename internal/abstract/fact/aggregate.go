package fact

type Aggregate interface {
	GetRemark() uint
	GetStory() uint
	PointToRemark(remark uint) error
	HasNextFact() (bool, error)
	ToNextFact() (Aggregate, error)
	GetFirstRemark() (uint, error)

	HasNextStory() (bool, error)
	ToNextStory() (Aggregate, error)

	HasParentStory() (bool, error)
	ToParentStory() (Aggregate, error)

	ToFirstFact() (Aggregate, error)
}
