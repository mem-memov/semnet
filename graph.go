package semnet

import (
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/detail"
	"github.com/mem-memov/semnet/internal/fact"
	"github.com/mem-memov/semnet/internal/phrase"
	"github.com/mem-memov/semnet/internal/remark"
	"github.com/mem-memov/semnet/internal/story"
	"github.com/mem-memov/semnet/internal/user"
	"github.com/mem-memov/semnet/internal/word"
)

type Graph struct {
	userRepository *user.Repository
}

func NewGraph(storage storage) *Graph {

	bitRepository := bit.NewRepository(storage)
	characterRepository := character.NewRepository(storage, bitRepository)
	wordRepository := word.NewRepository(storage, characterRepository)
	phraseRepository := phrase.NewRepository(storage, wordRepository)
	detailRepository := detail.NewRepository(storage, phraseRepository)
	remarkRepository := remark.NewRepository(storage, detailRepository)
	factRepository := fact.NewRepository(storage, remarkRepository)
	storyRepository := story.NewRepository(storage, factRepository)
	userRepository := user.NewRepository(storage, storyRepository)

	return &Graph{
		userRepository: userRepository,
	}
}

func (g *Graph) CreateUser() (User, error) {

	return User{}, nil
}

func (g *Graph) CreateStory(user User) (Story, error) {

	return Story{}, nil
}

func (g *Graph) CreateFact(story Story) (Fact, error) {

	return Fact{}, nil
}

func (g *Graph) CreateRemark(fact Fact, detail Detail) (Remark, error) {

	return Remark{}, nil
}

func (g *Graph) CreateDetail(object string, property string) (Detail, error) {

	return Detail{}, nil
}

func (g *Graph) GetDetailRemarks(detail Detail) ([]Remark, error) {

	return []Remark{}, nil
}

func (g *Graph) GetFactRemarks(fact Fact) ([]Remark, error) {

	return []Remark{}, nil
}

func (g *Graph) GetStoryFacts(story Story) ([]Fact, error) {

	return []Fact{}, nil
}

func (g *Graph) GetUserStories(user User) ([]Story, error) {

	return []Story{}, nil
}
