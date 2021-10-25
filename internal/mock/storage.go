package mock

type StorageMock struct {
	Has_ func(source uint) (bool, error)
	Create_ func() (uint, error)
	ReadSources_ func(target uint) ([]uint, error)
	ReadTargets_ func(source uint) ([]uint, error)
	SetReference_ func(source uint, reference uint) error
	GetReference_ func(source uint) (uint, uint, error)
	Connect_ func(source uint, target uint) error
	Disconnect_ func(source uint, target uint) error
	Delete_ func(source uint) error
}

func (s *StorageMock) Has(source uint) (bool, error) {
	return s.Has_(source)
}

func (s *StorageMock) Create() (uint, error) {
	return s.Create_()
}

func (s *StorageMock) ReadSources(target uint) ([]uint, error) {
	return s.ReadSources_(target)
}

func (s *StorageMock) ReadTargets(source uint) ([]uint, error) {
	return s.ReadTargets_(source)
}

func (s *StorageMock) SetReference(source uint, reference uint) error {
	return s.SetReference_(source, reference)
}

func (s *StorageMock) GetReference(source uint) (uint, uint, error) {
	return s.GetReference_(source)
}

func (s *StorageMock) Connect(source uint, target uint) error {
	return s.Connect_(source, target)
}

func (s *StorageMock) Disconnect(source uint, target uint) error {
	return s.Disconnect_(source, target)
}

func (s *StorageMock) Delete(source uint) error {
	return s.Delete_(source)
}
