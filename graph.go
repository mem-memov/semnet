package semnet

import (
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/detail"
	"github.com/mem-memov/semnet/internal/fact"
	"github.com/mem-memov/semnet/internal/phrase"
	"github.com/mem-memov/semnet/internal/remark"
	"github.com/mem-memov/semnet/internal/story"
	"github.com/mem-memov/semnet/internal/user"
	"github.com/mem-memov/semnet/internal/word"
)

type Graph struct {
	userRepository   *user.Repository
	storyRepository  *story.Repository
	factRepository   *fact.Repository
	remarkRepository *remark.Repository
	detailRepository *detail.Repository
	phraseRepository *phrase.Repository
}

func NewGraph(storage storage) *Graph {

	classRepository := class.NewRepository(storage)
	bitRepository := bit.NewRepository(storage, classRepository)
	characterRepository := character.NewRepository(storage, classRepository, bitRepository)
	wordRepository := word.NewRepository(storage, classRepository, characterRepository)
	phraseRepository := phrase.NewRepository(storage, classRepository, wordRepository)
	detailRepository := detail.NewRepository(storage, classRepository, phraseRepository)
	remarkRepository := remark.NewRepository(storage, classRepository, detailRepository)
	factRepository := fact.NewRepository(storage, classRepository, remarkRepository)
	storyRepository := story.NewRepository(storage, classRepository, factRepository)
	userRepository := user.NewRepository(storage, classRepository, storyRepository)

	return &Graph{
		userRepository:   userRepository,
		storyRepository:  storyRepository,
		factRepository:   factRepository,
		remarkRepository: remarkRepository,
		detailRepository: detailRepository,
		phraseRepository: phraseRepository,
	}
}

func (g *Graph) CreateUser() (User, error) {

	return User{}, nil
}

func (g *Graph) GetUser(userIdentifier uint) (User, error) {

	return User{}, nil
}

func (g *Graph) CreateRemark(
	userIdentifier uint,
	previousRemarkIdentifier uint,
	referencedRemarkIdentifiers []uint,
	isNewStory bool,
	isNewFact bool,
	object string,
	property string,
) (Remark, error) {

	userEntity, err := g.userRepository.GetUser(userIdentifier)
	if err != nil {
		return Remark{}, err
	}

	previousRemarkEntity, err := g.remarkRepository.GetRemark(previousRemarkIdentifier)
	if err != nil {
		return Remark{}, err
	}

	referencedRemarkEntities := make([]remark.Entity, len(referencedRemarkIdentifiers))
	for index, referencedRemarkIdentifier := range referencedRemarkIdentifiers {
		referencedRemarkEntity, err := g.remarkRepository.GetRemark(referencedRemarkIdentifier)
		if err != nil {
			return Remark{}, err
		}
		referencedRemarkEntities[index] = referencedRemarkEntity
	}

	var storyEntity story.Entity
	if isNewStory {
		storyEntity, err = g.storyRepository.CreateStory(userEntity.IdentifierForStory())
		if err != nil {
			return Remark{}, err
		}
	} else {
		storyEntity, err = previousRemarkEntity.GetStory(g.storyRepository)
		if err != nil {
			return Remark{}, err
		}
	}

	var factEntity fact.Entity
	if isNewFact {
		factEntity, err = g.factRepository.CreateFact(storyEntity.IdentifierForFact())
		if err != nil {
			return Remark{}, err
		}
	} else {
		factEntity, err = previousRemarkEntity.GetFact(g.storyRepository)
		if err != nil {
			return Remark{}, err
		}
	}

	remarkEntity, err := g.remarkRepository.CreateRemark(
		factEntity.IdentifierForRemark(),
		referencedRemarkEntities,
		object,
		property,
		)

	return Remark{
		id: remarkEntity.IdentifierForClass(),
		fact: Fact{
			id: factEntity.IdentifierForClass(),
			story: Story{
				id: storyEntity.IdentifierForClass(),
				user: userEntity.IdentifierForClass(),
			},
		},
		detail: Detail{
			object: object,
			property: property,
		},
	}, nil
}

func (g *Graph) GetRemark(remarkIdentifier uint) (Remark, error) {

	return Remark{}, nil
}
