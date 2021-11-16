package story

type Entity interface {
	GetClass() uint
	GetFact() uint
	GetPosition() uint
	GetTree() uint

	PointToFact(fact uint) error
	PointToPosition(position uint) error
	PointToTree(tree uint) error

	HasTargetPosition() (bool, error)
	GetTargetPosition() (uint, error)

	GetTargetFact() (uint, error)

	HasTargetTree() (bool, error)
	GetTargetTree() (uint, error)

	HasSourceTree() (bool, error)
	GetSourceTree() (uint, error)
}
