package detail

import (
	"fmt"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Factory struct {
	storage abstractDetail.Storage
}

var _ abstractDetail.Factory = &Factory{}

func NewFactory(storage abstractDetail.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) ProvideEntity(
	classEntity abstractClass.Entity,
	objectPhrase abstractPhrase.Aggregate,
	propertyPhrase abstractPhrase.Aggregate,
) (abstractDetail.Entity, error) {

	objectTargetDetails, err := objectPhrase.GetTargetDetails()
	if err != nil {
		return nil, err
	}

	propertySourceDetails, err := propertyPhrase.GetSourceDetails()

	commonDetailIdentifiers := make([]uint, 0, 1)

	// TODO: optimize (cut tails, use map, sort)
	for _, objectTargetDetail := range objectTargetDetails {
		for _, propertySourceDetail := range propertySourceDetails {
			if propertySourceDetail == objectTargetDetail {
				commonDetailIdentifiers = append(commonDetailIdentifiers, propertySourceDetail)
			}
		}
	}

	switch len(commonDetailIdentifiers) {
	case 0:
		return f.storage.CreateEntity(classEntity, objectPhrase, propertyPhrase)
	case 1:
		return f.storage.ReadEntityByPhrase(commonDetailIdentifiers[0])
	default:
		return Entity{}, fmt.Errorf("phrases have too many common details")
	}
}
