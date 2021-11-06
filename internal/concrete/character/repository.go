package character

import (
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	"github.com/mem-memov/semnet/internal/concrete/bit"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Repository struct {
	entities      *entities
	bitRepository *bit.Repository
	tree          *tree
	paths         *paths
}

func NewRepository(storage storage, classRepository *class.Repository, bitRepository *bit.Repository) *Repository {
	entities := newEntities(storage, classRepository, bitRepository)

	return &Repository{
		entities:      entities,
		bitRepository: bitRepository,
		tree:          newTree(storage, entities),
		paths:         newPaths(),
	}
}

func (r *Repository) Provide(integer rune) (interface{}, error) {

	path, err := r.paths.collect(integer)
	if err != nil {
		return Entity{}, err
	}

	firstBit, err := r.bitRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.provideRoot(firstBit)
	if err != nil {
		return Entity{}, err
	}

	for _, bitValue := range path[1:] {

		entity, err = entity.provideNext(bitValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

func (r *Repository) Extract(entity interface{}) (rune, error) {

	entity_ := entity.(abstractCharacter.Entity) // TODO: remove

	bitValue, err := entity_.BitValue()
	if err != nil {
		return 0, err
	}

	path := r.paths.create(bitValue)

	for {
		var isRoot bool
		entity_, isRoot, err = entity_.FindPrevious(r.entities)

		if isRoot {
			break
		}

		bitValue, err = entity_.BitValue()
		if err != nil {
			return 0, err
		}

		path = append(path, bitValue)
	}

	var integer int32

	for _, bitValue := range path.reverse() {
		integer <<= 1
		if bitValue {
			integer += 1
		}
	}

	return integer, nil
}

func (r *Repository) Fetch(wordIdentifier uint) (interface{}, error) {

	return r.entities.createWithWord(wordIdentifier)
}
