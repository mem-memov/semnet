package class

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/mock"
	classMock "github.com/mem-memov/semnet/internal/mock/class"
	"github.com/mem-memov/semnet/internal/mock/fact/class"
	"testing"
)

func TestNode_IsValid_ReturnsTrue(t *testing.T) {
	entityInstance := classMock.EntityMock{
		IsFact_: func(uint) (bool, error) {
			return true, nil
		},
	}

	repositoryInstance := &classMock.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return entityInstance, nil
		},
	}

	creatorInstance := &class.IdentifierCreatorMock{}

	node := newNode(1, repositoryInstance, creatorInstance)

	isValid, err := node.IsValid()
	if err != nil {
		t.Fail()
	}

	if !isValid {
		t.Fail()
	}
}

func TestNode_IsValid_ReturnsFalse(t *testing.T) {
	entityInstance := classMock.EntityMock{
		IsFact_: func(uint) (bool, error) {
			return false, nil
		},
	}

	repositoryInstance := &classMock.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return entityInstance, nil
		},
	}

	creatorInstance := &class.IdentifierCreatorMock{}

	node := newNode(1, repositoryInstance, creatorInstance)

	isValid, err := node.IsValid()
	if err != nil {
		t.Fail()
	}

	if isValid {
		t.Fail()
	}
}

func TestNode_IsValid_RepositoryFailsProvidingEntity(t *testing.T) {
	errorInstance := mock.ErrorMock{
		Error_: func() string {
			return "repository failure"
		},
	}

	repositoryInstance := &classMock.RepositoryMock{
		ProvideEntity_: func() (abstractClass.Entity, error) {
			return nil, errorInstance
		},
	}

	creatorInstance := &class.IdentifierCreatorMock{}

	node := newNode(1, repositoryInstance, creatorInstance)

	_, err := node.IsValid()
	if err == nil {
		t.Fail()
	}
}

func TestNode_CreateNewNode_Success(t *testing.T) {
	repositoryInstance := &classMock.RepositoryMock{}

	creatorInstance := &class.IdentifierCreatorMock{
		CreateNewIdentifier_: func() (uint, error) {
			return 2, nil
		},
	}

	node := newNode(1, repositoryInstance, creatorInstance)

	_, err := node.CreateNewNode()
	if err != nil {
		t.Fail()
	}
}

func TestNode_CreateNewNode_CreatorFails(t *testing.T) {
	repositoryInstance := &classMock.RepositoryMock{}

	errorInstance := &mock.ErrorMock{
		Error_: func() string {
			return "creator failure"
		},
	}

	creatorInstance := &class.IdentifierCreatorMock{
		CreateNewIdentifier_: func() (uint, error) {
			return 0, errorInstance
		},
	}

	node := newNode(1, repositoryInstance, creatorInstance)

	_, err := node.CreateNewNode()
	if err == nil {
		t.Fail()
	}
}
