package semnet

import (
	"github.com/mem-memov/semnet/internal/abstract"
	"github.com/mem-memov/semnet/internal/concrete/bit"
	"github.com/mem-memov/semnet/internal/concrete/character"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/detail"
	"github.com/mem-memov/semnet/internal/concrete/fact"
	"github.com/mem-memov/semnet/internal/concrete/phrase"
	"github.com/mem-memov/semnet/internal/concrete/remark"
	"github.com/mem-memov/semnet/internal/concrete/story"
	"github.com/mem-memov/semnet/internal/concrete/word"
)

type Graph struct {
	storyRepository  *story.Repository
	factRepository   *fact.Repository
	remarkRepository *remark.Repository
	detailRepository *detail.Repository
	phraseRepository *phrase.Repository
}

func NewGraph(storage abstract.Storage) *Graph {

	classRepository := class.NewRepository(storage)
	bitRepository := bit.NewRepository(storage, classRepository)
	characterRepository := character.NewRepository(storage, classRepository, bitRepository)
	wordRepository := word.NewRepository(storage, classRepository, characterRepository)
	phraseRepository := phrase.NewRepository(storage, classRepository, wordRepository)
	detailRepository := detail.NewRepository(storage, classRepository, phraseRepository)

	storyRepository := story.NewRepository(storage, classRepository)
	factRepository := fact.NewRepository(storage, classRepository, storyRepository)

	remarkRepository := remark.NewRepository(storage, classRepository, detailRepository, factRepository)


	return &Graph{
		remarkRepository: remarkRepository,
	}
}

func (g *Graph) CreateStory(object string, property string) (Remark, error) {
	return g.remarkRepository.CreateFirstUserRemark(object, property)
}

func (g *Graph) GetRemark(remarkIdentifier uint) (Remark, error) {
	return g.remarkRepository.GetRemark(remarkIdentifier)
}
