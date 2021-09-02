package remark

type storage interface {
	Has(source uint) (bool, error)
	Create() (uint, error)
	ReadSources(target uint) ([]uint, error)
	ReadTargets(source uint) ([]uint, error)
	SetReference(source uint, reference uint) error
	GetReference(source uint) (uint, uint, error)
	Connect(source uint, target uint) error
	Disconnect(source uint, target uint) error
	Delete(source uint) error
}
