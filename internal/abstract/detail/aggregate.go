package detail

type Aggregate interface {
	PointToRemark(remark uint) error
	GetObjectAndProperty() (string, string, error)
}
