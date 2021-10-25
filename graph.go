package semnet

import (
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

	return &Graph{
		storyRepository:  storyRepository,
		factRepository:   factRepository,
		remarkRepository: remarkRepository,
		detailRepository: detailRepository,
		phraseRepository: phraseRepository,
	}
}

func (g *Graph) CreateStory() (Story, error) {
	return g.storyRepository.CreateStory()
}

func (g *Graph) GetRemark(remarkIdentifier uint) (Remark, error) {
	return g.remarkRepository.GetRemark(remarkIdentifier)
}
