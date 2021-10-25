package node

type Classes interface {
	Create(identifier uint) Class
	CreateNew() (Class, error)
}
