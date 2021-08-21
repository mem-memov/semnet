package node

type Codes struct {
	storage storage
}

func NewCodes(storage storage) *Codes {
	return &Codes{
		storage: storage,
	}
}

func (c *Codes) Create(identifier uint) Code {
	return newCode(identifier, c.storage)
}
