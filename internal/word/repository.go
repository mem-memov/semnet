package word

import (
	"github.com/mem-memov/semnet/internal/character"
)

type Repository struct {
	entities            *entities
	characterRepository *character.Repository
	tree                *tree
	paths               *paths
}

func NewRepository(storage storage, characterRepository *character.Repository) *Repository {
	entities := newEntities(storage, characterRepository)

	return &Repository{
		entities:            entities,
		characterRepository: characterRepository,
		tree:                newLayer(storage, entities),
		paths:               newPaths(),
	}
}

func (r *Repository) Provide(word string) (Entity, error) {

	path, err := r.paths.collect(word)
	if err != nil {
		return Entity{}, err
	}

	firstCharacter, err := r.characterRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.provideRoot(firstCharacter)
	if err != nil {
		return Entity{}, err
	}

	for _, characterValue := range path[1:] {

		entity, err = entity.provideNext(characterValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

//func (r *Repository) create(rune rune) (Entity, error) {
//
//	character, err := c.characters.create(int32(rune))
//	if err != nil {
//		return Entity{}, err
//	}
//
//	var character character
//	var err error
//
//	for i, bitName := range fmt.Sprintf("%b", r) {
//
//		switch bitName {
//		case '0':
//			if i == 0 {
//				character, err = c.characters.createZero()
//				if err != nil {
//					return Entity{}, err
//				}
//			} else {
//				character, err = character.NextZero()
//				if err != nil {
//					return Entity{}, err
//				}
//			}
//		case '1':
//			if i == 0 {
//				character, err = c.characters.createOne()
//				if err != nil {
//					return Entity{}, err
//				}
//			} else {
//				character, err = character.NextOne()
//				if err != nil {
//					return Entity{}, err
//				}
//			}
//		default:
//			return Entity{}, fmt.Errorf("unexpected bit name: %c", bitName)
//		}
//	}
//
//	return newCharacter(character), nil
//}
