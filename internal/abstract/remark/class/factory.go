package class

type Factory interface {
	CreateExistingNode(identifier uint) Node
	CreateNewNode() (Node, error)
}
