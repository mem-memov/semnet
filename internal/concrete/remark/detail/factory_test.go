package detail

import (
	"github.com/mem-memov/semnet/internal/mock"
	detailMock "github.com/mem-memov/semnet/internal/mock/detail"
	creatorMock "github.com/mem-memov/semnet/internal/mock/remark/detail"
	"testing"
)

func TestFactory_CreateExistingNode(t *testing.T) {
	storageInstance := &mock.StorageMock{}
	repositoryInstance := &detailMock.RepositoryMock{}
	creatorInstance := &creatorMock.IdentifierCreatorMock{}

	factory := NewFactory(storageInstance, repositoryInstance, creatorInstance)

	node := factory.CreateExistingNode(1)

	if node.identifier != 1 {
		t.Fail()
	}
}

func TestFactory_CreateNewNode_Success(t *testing.T) {
	storageInstance := &mock.StorageMock{}
	repositoryInstance := &detailMock.RepositoryMock{}

	creatorInstance := &creatorMock.IdentifierCreatorMock{
		CreateNewIdentifier_: func(object string, property string) (uint, error) {
			return 2, nil
		},
	}

	factory := NewFactory(storageInstance, repositoryInstance, creatorInstance)

	newNode, err := factory.CreateNewNode("sings", "a bird")
	if err != nil {
		t.Fail()
	}

	if newNode.identifier != 2 {
		t.Fail()
	}
}

func TestFactory_CreateNewNode_CreatorFails(t *testing.T) {
	storageInstance := &mock.StorageMock{}
	repositoryInstance := &detailMock.RepositoryMock{}

	errorInstance := &mock.ErrorMock{
		Error_: func() string {
			return "creator failure"
		},
	}

	creatorInstance := &creatorMock.IdentifierCreatorMock{
		CreateNewIdentifier_: func(object string, property string) (uint, error) {
			return 0, errorInstance
		},
	}

	factory := NewFactory(storageInstance, repositoryInstance, creatorInstance)

	_, err := factory.CreateNewNode("sings", "a bird")
	if err == nil {
		t.Fail()
	}

	if err.Error() != "creator failure" {
		t.Fail()
	}
}
