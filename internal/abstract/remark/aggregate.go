package remark

type Aggregate interface {
	GetIdentifier() uint
	HasNextRemark() (bool, error)
	GetNextRemark() (Aggregate, error)
	HasNextFact() (bool, error)
	GetNextFact() (Aggregate, error)
	GetFirstFact() (Aggregate, error)
	HasNextStory() (bool, error)
	GetNextStory() (Aggregate, error)
	GetObjectAndProperty() (string, string, error)
	AddRemarkToFact(property string) (Aggregate, error)
	AddFactToStory(object string, property string) (Aggregate, error)
	CreateChildStory(object string, property string) (Aggregate, error)
}
