package class

import abstractClass "github.com/mem-memov/semnet/internal/abstract/class"

type RepositoryMock struct {
	ProvideEntity_ func() (abstractClass.Entity, error)
}

var _ abstractClass.Repository = &RepositoryMock{}

func (r *RepositoryMock) ProvideEntity() (abstractClass.Entity, error) {
	return r.ProvideEntity_()
}
