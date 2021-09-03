package node

type Remarks struct {
	storage storage
}

func NewRemarks(storage storage) *Remarks {
	return &Remarks{
		storage: storage,
	}
}

func (r *Remarks) Create(identifier uint) Remark {
	return newRemark(identifier, r.storage)
}
