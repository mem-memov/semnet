package class

type Node interface {
	GetIdentifier() uint
	IsValid() (bool, error)
	CreateNewNode() (Node, error)
}
