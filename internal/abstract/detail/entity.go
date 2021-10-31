package detail

type Entity interface {
	GetClass() uint
	GetPhrase() uint
	GetRemark() uint
	PointToRemark(remark uint) error
	GetObjectAndPropertyPhrases() (uint, uint, error)
}
