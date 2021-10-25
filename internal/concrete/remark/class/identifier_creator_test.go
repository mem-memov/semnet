package class

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/mock"
	mockClass "github.com/mem-memov/semnet/internal/mock/class"
	"testing"
)

func TestCreator_CreateNewNode_Successful(t *testing.T) {
	entityInstance := mockClass.EntityMock{
		CreateRemark_: func() (uint, error) {
			return 1, nil
		},
	}

	repositoryInstance := &mockClass.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return entityInstance, nil
		},
	}

	creator := newIdentifierCreator(repositoryInstance)

	identifier, err := creator.CreateNewIdentifier()
	if err != nil {
		t.Fail()
	}

	if identifier != 1 {
		t.Fail()
	}
}

func TestCreator_CreateNewNode_RepositoryFailsProvidingEntity(t *testing.T) {
	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "repository failure"
		},
	}

	repositoryInstance := &mockClass.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return nil, errorInstance
		},
	}

	creator := newIdentifierCreator(repositoryInstance)

	_, err := creator.CreateNewIdentifier()
	if err == nil {
		t.Fail()
	}

	if err.Error() != "repository failure" {
		t.Fail()
	}
}

func TestCreator_CreateNewNode_ClassEntityFailsCreatingRemark(t *testing.T) {
	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "cluster failure"
		},
	}

	entityInstance := mockClass.EntityMock{
		CreateRemark_: func() (uint, error) {
			return 0, errorInstance
		},
	}

	repositoryInstance := &mockClass.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return entityInstance, nil
		},
	}

	creator := newIdentifierCreator(repositoryInstance)

	_, err := creator.CreateNewIdentifier()
	if err == nil {
		t.Fail()
	}

	if err.Error() != "cluster failure" {
		t.Fail()
	}
}
