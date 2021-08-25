package node

type Details struct {
	storage storage
}

func NewDetails(storage storage) *Details {
	return &Details{
		storage: storage,
	}
}

func (d *Details) Create(identifier uint) Detail {
	return newDetail(identifier, d.storage)
}