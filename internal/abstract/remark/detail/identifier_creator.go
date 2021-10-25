package detail

type IdentifierCreator interface {
	CreateNewIdentifier(object string, property string) (uint, error)
}
