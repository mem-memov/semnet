package detail

import (
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	"github.com/mem-memov/semnet/internal/mock"
	mockDetail "github.com/mem-memov/semnet/internal/mock/detail"
	"testing"
)

func TestCreator_CreateNewNode_Successful(t *testing.T) {
	storageInstance := &mock.StorageMock{
		Create_: func() (uint, error) {
			return 1, nil
		},
	}

	entityInstance := mockDetail.EntityMock{
		AddRemark_: func(remarkIdentifier uint) error {
			return nil
		},
	}

	repositoryInstance := &mockDetail.RepositoryMock{
		Provide_: func(object string, property string) (abstractDetail.Entity, error) {
			return entityInstance, nil
		},
	}

	creator := newIdentifierCreator(storageInstance, repositoryInstance)

	identifier, err := creator.CreateNewIdentifier("sings", "a bird")
	if err != nil {
		t.Fail()
	}

	if identifier != 1 {
		t.Fail()
	}
}

func TestCreator_CreateNewNode_RepositoryFailsProvidingEntity(t *testing.T) {

	storageInstance := &mock.StorageMock{}

	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "repository failure"
		},
	}

	repositoryInstance := &mockDetail.RepositoryMock{
		Provide_: func(object string, property string) (abstractDetail.Entity, error) {
			return nil, errorInstance
		},
	}

	creator := newIdentifierCreator(storageInstance, repositoryInstance)

	_, err := creator.CreateNewIdentifier("sings", "a bird")
	if err == nil {
		t.Fail()
	}

	if err.Error() != "repository failure" {
		t.Fail()
	}
}

func TestCreator_CreateNewNode_StorageFailsCreatingIdentifier(t *testing.T) {
	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "storage failure"
		},
	}

	storageInstance := &mock.StorageMock{
		Create_: func() (uint, error) {
			return 0, errorInstance
		},
	}

	entityInstance := mockDetail.EntityMock{}

	repositoryInstance := &mockDetail.RepositoryMock{
		Provide_: func(object string, property string) (abstractDetail.Entity, error) {
			return entityInstance, nil
		},
	}

	creator := newIdentifierCreator(storageInstance, repositoryInstance)

	_, err := creator.CreateNewIdentifier("sings", "a bird")
	if err == nil {
		t.Fail()
	}

	if err.Error() != "storage failure" {
		t.Fail()
	}
}

func TestCreator_CreateNewNode_DetailEntityFailsAddingRemark(t *testing.T) {
	storageInstance := &mock.StorageMock{
		Create_: func() (uint, error) {
			return 1, nil
		},
	}

	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "entity failure"
		},
	}

	entityInstance := mockDetail.EntityMock{
		AddRemark_: func(remarkIdentifier uint) error {
			return errorInstance
		},
	}

	repositoryInstance := &mockDetail.RepositoryMock{
		Provide_: func(object string, property string) (abstractDetail.Entity, error) {
			return entityInstance, nil
		},
	}

	creator := newIdentifierCreator(storageInstance, repositoryInstance)

	_, err := creator.CreateNewIdentifier("sings", "a bird")
	if err == nil {
		t.Fail()
	}

	if err.Error() != "entity failure" {
		t.Fail()
	}
}