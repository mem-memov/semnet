package semnet

type Remark struct {
	id     uint
	fact   Fact
	detail Detail
}

func newRemark(id uint, fact Fact, detail Detail) Remark {
	return Remark{
		id:     id,
		fact:   fact,
		detail: detail,
	}
}

func (r Remark) GetId() uint {
	return r.id
}

func (r Remark) GetFact() Fact {
	return r.fact
}

func (r Remark) GetDetail() Detail {
	return r.detail
}
