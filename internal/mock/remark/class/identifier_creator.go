package class

type IdentifierCreatorMock struct {
	CreateNewIdentifier_ func() (uint, error)
}

func (c IdentifierCreatorMock) CreateNewIdentifier() (uint, error) {
	return c.CreateNewIdentifier_()
}