package detail

type Factory interface {
	CreateExistingNode(identifier uint) Node
	CreateNewNode(object string, property string) (Node, error)
}
