package node

type Remark struct {
	identifier uint
}

func NewRemark(identifier uint) Remark {
	return Remark{
		identifier: identifier,
	}
}

func (r Remark) Identifier() uint {
	return r.identifier
}
