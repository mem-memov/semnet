package class

import (
	"github.com/mem-memov/semnet/internal/mock"
	classMock "github.com/mem-memov/semnet/internal/mock/class"
	"github.com/mem-memov/semnet/internal/mock/story/class"
	"testing"
)

func TestFactory_CreateNew(t *testing.T) {
	repositoryInstance := &classMock.RepositoryMock{}

	creatorInstance := &class.IdentifierCreatorMock{}

	factory := &Factory{
		repository: repositoryInstance,
		creator: creatorInstance,
	}

	_ = factory.CreateExistingNode(1)
}

func TestFactory_CreateNewNode_Success(t *testing.T) {
	repositoryInstance := &classMock.RepositoryMock{}

	creatorInstance := &class.IdentifierCreatorMock{
		CreateNewIdentifier_: func() (uint, error) {
			return 2, nil
		},
	}

	factory := &Factory{
		repository: repositoryInstance,
		creator: creatorInstance,
	}

	_, err := factory.CreateNewNode()
	if err != nil {
		t.Fail()
	}
}

func TestFactory_CreateNewNode_CreatorFails(t *testing.T) {
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

	factory := &Factory{
		repository: repositoryInstance,
		creator: creatorInstance,
	}

	_, err := factory.CreateNewNode()
	if err == nil {
		t.Fail()
	}

	if err.Error() != "creator failure" {
		t.Fail()
	}
}
