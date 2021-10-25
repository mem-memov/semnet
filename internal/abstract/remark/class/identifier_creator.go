package class

type IdentifierCreator interface {
	CreateNewIdentifier() (uint, error)
}