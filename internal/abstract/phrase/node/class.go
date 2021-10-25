package node

type Class interface {
	IsValid() (bool, error)
	NewClass() (Class, error)
}
