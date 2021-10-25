package class

type Node interface {
	IsValid() (bool, error)
	CreateNewNode() (Node, error)
}
