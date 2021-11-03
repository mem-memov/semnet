package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Repository struct {
	detailStorage    abstractDetail.Storage
	detailFactory    abstractDetail.Factory
	classRepository  abstractClass.Repository
	phraseRepository abstractPhrase.Repository
}

var _ abstractDetail.Repository = &Repository{}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	phraseRepository abstractPhrase.Repository,
) *Repository {

	detailStorage := NewStorage(storage)

	return &Repository{
		detailStorage:    detailStorage,
		detailFactory:    NewFactory(detailStorage),
		classRepository:  classRepository,
		phraseRepository: phraseRepository,
	}
}

func (r *Repository) Provide(object string, property string) (abstractDetail.Aggregate, error) {

	objectPhrase, err := r.phraseRepository.Provide(object)
	if err != nil {
		return nil, err
	}

	propertyPhrase, err := r.phraseRepository.Provide(property)
	if err != nil {
		return nil, err
	}

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	detail, err := r.detailFactory.ProvideEntity(class, objectPhrase, propertyPhrase)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		detail:           detail,
		detailStorage:    r.detailStorage,
		detailFactory:    r.detailFactory,
		classRepository:  r.classRepository,
		phraseRepository: r.phraseRepository,
	}, nil
}

func (r *Repository) Fetch(remark uint) (abstractDetail.Aggregate, error) {

	detail, err := r.detailStorage.ReadEntityByRemark(remark)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		detail:           detail,
		detailStorage:    r.detailStorage,
		detailFactory:    r.detailFactory,
		classRepository:  r.classRepository,
		phraseRepository: r.phraseRepository,
	}, nil
}
