package detail

import (
	"errors"
	"github.com/mem-memov/semnet/internal/mock"
	mockClass "github.com/mem-memov/semnet/internal/mock/class"
	mockPhrase "github.com/mem-memov/semnet/internal/mock/phrase"
	"testing"
)

func TestStorage_CreateEntity_Successfully(t *testing.T) {

	identifier := uint(1000)
	storageMock := &mock.StorageMock{
		Create_: func() (uint, error) {
			identifier++
			return identifier, nil
		},
		SetReference_: func(source uint, reference uint) error {
			if (source != 555 || reference != 1001) && (source != 1001 || reference != 1002) {
				t.Fail()
			}
			return nil
		},
	}

	classMock := mockClass.EntityMock{
		CreateDetail_: func() (uint, error) {
			return 555, nil
		},
	}

	objectPhraseMock := mockPhrase.EntityMock{
		AddTargetDetail_: func(detail uint) error {
			if detail != 1001 {
				t.Fail()
			}
			return nil
		},
	}

	propertyPhraseMock := mockPhrase.EntityMock{
		AddSourceDetail_: func(detail uint) error {
			if detail != 1001 {
				t.Fail()
			}
			return nil
		},
	}

	storage := &Storage{
		storage: storageMock,
	}

	entity, err := storage.CreateEntity(classMock, objectPhraseMock, propertyPhraseMock)
	if err != nil {
		t.Fail()
	}

	if entity.GetClass() != 555 {
		t.Fail()
	}

	if entity.GetPhrase() != 1001 {
		t.Fail()
	}

	if entity.GetRemark() != 1002 {
		t.Fail()
	}
}

func TestStorage_CreateEntity_FailSettingReferenceFromPhraseToRemark(t *testing.T) {

	identifier := uint(1000)
	storageMock := &mock.StorageMock{
		Create_: func() (uint, error) {
			identifier++
			return identifier, nil
		},
		SetReference_: func(source uint, reference uint) error {
			if (source != 555 || reference != 1001) && (source != 1001 || reference != 1002) {
				t.Fail()
			}

			if source == 1001 && reference == 1002 {
				return errors.New("FailSettingReferenceFromPhraseToRemark")
			}
			return nil
		},
	}

	classMock := mockClass.EntityMock{
		CreateDetail_: func() (uint, error) {
			return 555, nil
		},
	}

	objectPhraseMock := mockPhrase.EntityMock{
		AddTargetDetail_: func(detail uint) error {
			if detail != 1001 {
				t.Fail()
			}
			return nil
		},
	}

	propertyPhraseMock := mockPhrase.EntityMock{
		AddSourceDetail_: func(detail uint) error {
			if detail != 1001 {
				t.Fail()
			}
			return nil
		},
	}

	storage := &Storage{
		storage: storageMock,
	}

	_, err := storage.CreateEntity(classMock, objectPhraseMock, propertyPhraseMock)
	if err == nil {
		t.Fail()
	}

	if err.Error() != "FailSettingReferenceFromPhraseToRemark" {
		t.Fail()
	}
}
