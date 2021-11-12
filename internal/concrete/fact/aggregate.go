package fact

import (
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Aggregate struct {
	fact            abstractFact.Entity
	factStorage     abstractFact.Storage
	storyRepository abstractStory.Repository
}

var _ abstractFact.Aggregate = Aggregate{}

func (a Aggregate) GetRemark() uint {

	return a.fact.GetRemark()
}

func (a Aggregate) GetStory() uint {

	return a.fact.GetStory()
}

func (a Aggregate) PointToRemark(remark uint) error {

	return a.fact.PointToRemark(remark)
}

func (a Aggregate) HasNextFact() (bool, error) {

	return a.fact.HasTargetFact()
}

func (a Aggregate) ToNextFact() (abstractFact.Aggregate, error) {

	nextFactIdentifier, err := a.fact.GetTargetFact()
	if err != nil {
		return nil, err
	}

	nextFact, err := a.factStorage.ReadEntityByPosition(nextFactIdentifier)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		fact:            nextFact,
		storyRepository: a.storyRepository,
	}, nil
}

func (a Aggregate) GetFirstRemark() (uint, error) {

	return a.fact.GetFirstRemark()
}

func (a Aggregate) HasNextStory() (bool, error) {

	storyIdentifier, err := a.fact.GetTargetStory()

	story, err := a.storyRepository.FetchByFact(storyIdentifier)
	if err != nil {
		return false, err
	}

	return story.HasNextStory()
}

func (a Aggregate) ToNextStory() (abstractFact.Aggregate, error) {

	storyIdentifier, err := a.fact.GetTargetStory()

	story, err := a.storyRepository.FetchByFact(storyIdentifier)
	if err != nil {
		return nil, err
	}

	nextStoryIdentifier, err := story.GetTargetStory()
	if err != nil {
		return nil, err
	}

	nextStory, err := a.storyRepository.FetchByPosition(nextStoryIdentifier)
	if err != nil {
		return nil, err
	}

	nextFactIdentifier, err := nextStory.GetTargetFact()
	if err != nil {
		return nil, err
	}

	nextFact, err := a.factStorage.ReadEntityByStory(nextFactIdentifier)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		fact:            nextFact,
		storyRepository: a.storyRepository,
	}, nil
}
