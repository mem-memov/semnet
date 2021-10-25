package detail

import "github.com/mem-memov/semnet/internal/abstract/detail"

type RepositoryMock struct {
	Extend_ func(objectIdentifier uint, property string) (detail.Entity, error)
	Provide_ func(object string, property string) (detail.Entity, error)
	Fetch_ func(remarkIdentifier uint) (detail.Entity, error)
}

func (r *RepositoryMock) Extend(objectIdentifier uint, property string) (detail.Entity, error) {
	return r.Extend_(objectIdentifier, property)
}

func (r *RepositoryMock) Provide(object string, property string) (detail.Entity, error) {
	return r.Provide_(object, property)
}

func (r *RepositoryMock) Fetch(remarkIdentifier uint) (detail.Entity, error) {
	return r.Fetch_(remarkIdentifier)
}