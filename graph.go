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
	userRepository *user.Repository
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
		userRepository: userRepository,
	}
}

func (g *Graph) CreateUser() (User, error) {

	return User{}, nil
}

func (g *Graph) GetUser(userId uint) (User, error) {

	return User{}, nil
}
