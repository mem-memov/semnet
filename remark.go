package semnet

type Remark interface {
	GetFact() uint
	HasNextRemark() (bool, error)
	GetNextRemark() (Remark, error)
	HasNextFact() (bool, error)
	GetNextFact() (Remark, error)
	AddRemarkToFact(property string) (Remark, error)
}
