package semnet

type Remark interface {
	AddRemarkToFact(property string) (Remark, error)
}
