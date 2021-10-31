package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Repository struct {
	storage          abstract.Storage
	classRepository  abstractClass.Repository
	phraseRepository abstractPhrase.Repository
}

var _ abstractDetail.Repository = &Repository{}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	phraseRepository abstractPhrase.Repository,
) *Repository {

	return &Repository{
		storage:          storage,
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

	detail, err := createEntity(r.storage, class, objectPhrase, propertyPhrase)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		detail:           detail,
		storage:          r.storage,
		classRepository:  r.classRepository,
		phraseRepository: r.phraseRepository,
	}, nil
}

func (r *Repository) Fetch(remark uint) (abstractDetail.Aggregate, error) {

	detail, err := readEntityByRemark(r.storage, remark)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		detail:           detail,
		storage:          r.storage,
		classRepository:  r.classRepository,
		phraseRepository: r.phraseRepository,
	}, nil
}
