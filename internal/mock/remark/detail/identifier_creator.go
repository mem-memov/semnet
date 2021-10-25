package detail

type IdentifierCreatorMock struct {
	CreateNewIdentifier_ func(object string, property string) (uint, error)
}

func (i *IdentifierCreatorMock) CreateNewIdentifier(object string, property string) (uint, error) {
	return i.CreateNewIdentifier_(object, property)
}