package character

import (
	"github.com/mem-memov/semnet/internal/code"
)

type Repository struct {
	entities       *entities
	codeRepository *code.Repository
	tree           *tree
	paths          *paths
}

func NewRepository(storage storage, codeRepository *code.Repository) *Repository {
	entities := newEntities(storage, codeRepository)

	return &Repository{
		entities:       entities,
		codeRepository: codeRepository,
		tree:           newLayer(storage, entities),
		paths:          newPaths(),
	}
}

func (r *Repository) Provide(word string) (Entity, error) {

	path, err := r.paths.collect(word)
	if err != nil {
		return Entity{}, err
	}

	firstCode, err := r.codeRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.provideRoot(firstCode)
	if err != nil {
		return Entity{}, err
	}

	for _, codeValue := range path[1:] {

		entity, err = entity.provideNext(codeValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

//func (r *Repository) create(rune rune) (Entity, error) {
//
//	code, err := c.codes.create(int32(rune))
//	if err != nil {
//		return Entity{}, err
//	}
//
//	var code code
//	var err error
//
//	for i, bitName := range fmt.Sprintf("%b", r) {
//
//		switch bitName {
//		case '0':
//			if i == 0 {
//				code, err = c.codes.createZero()
//				if err != nil {
//					return Entity{}, err
//				}
//			} else {
//				code, err = code.NextZero()
//				if err != nil {
//					return Entity{}, err
//				}
//			}
//		case '1':
//			if i == 0 {
//				code, err = c.codes.createOne()
//				if err != nil {
//					return Entity{}, err
//				}
//			} else {
//				code, err = code.NextOne()
//				if err != nil {
//					return Entity{}, err
//				}
//			}
//		default:
//			return Entity{}, fmt.Errorf("unexpected bit name: %c", bitName)
//		}
//	}
//
//	return newCharacter(code), nil
//}
