package node

type Remark struct {
	identifier uint
	storage    storage
}

func newRemark(identifier uint, storage storage) Remark {
	return Remark{
		identifier: identifier,
		storage:    storage,
	}
}

func (r Remark) NewRemark(detail Detail) (Remark, error) {

	identifier, err := r.storage.Create()
	if err != nil {
		return Remark{}, nil
	}

	err = r.storage.SetReference(detail.Identifier(), identifier)
	if err != nil {
		return Remark{}, nil
	}

	return newRemark(identifier, r.storage), nil
}
