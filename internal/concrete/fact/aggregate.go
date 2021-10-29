package fact

import (
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Aggregate struct {
	entity          abstractFact.Entity
	storyRepository abstractStory.Repository
}

var _ abstractFact.Aggregate = Aggregate{}

func (a Aggregate) GetRemark() uint {

	return a.entity.GetRemark()
}

func (a Aggregate) GetStory() uint {

	return a.entity.GetStory()
}

func (a Aggregate) PointToRemark(remark uint) error {

	return a.entity.PointToRemark(remark)
}

func (a Aggregate) HasNextFact() (bool, error) {

	return a.entity.HasNextFact()
}

func (a Aggregate) ToNextFact() (abstractFact.Aggregate, error) {

	nextEntity, err := a.entity.GetNextFact()
	if err != nil {
		return nil, err
	}

	return Aggregate{
		entity:          nextEntity,
		storyRepository: a.storyRepository,
	}, nil
}

func (a Aggregate) GetFirstRemark() (uint, error) {

	return a.entity.GetFirstRemark()
}

func (a Aggregate) HasNextStory() (bool, error) {

	storyIdentifier, err := a.entity.GetTargetStory()

	story, err := a.storyRepository.FetchByFact(storyIdentifier)
	if err != nil {
		return false, err
	}

	return story.HasNextStory()
}

func (a Aggregate) ToNextStory() (abstractFact.Aggregate, error) {

	storyIdentifier, err := a.entity.GetTargetStory()

	story, err := a.storyRepository.FetchByFact(storyIdentifier)
	if err != nil {
		return nil, err
	}

	nextStory, err := story.GetNextStory()
	if err != nil {
		return nil, err
	}

	nextFactIdentifier, err := nextStory.GetTargetFact()
	if err != nil {
		return nil, err
	}

	nextFact, err := a.entity.ToNextStory(nextFactIdentifier)

	return Aggregate{
		entity:          nextFact,
		storyRepository: a.storyRepository,
	}, nil
}
