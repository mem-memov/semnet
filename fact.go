package semnet

type Fact struct {
	id    uint
	story Story
}

func newFact(id uint, story Story) Fact {
	return Fact{
		id:    id,
		story: story,
	}
}

func (f Fact) GetId() uint {
	return f.id
}

func (f Fact) GetStory() Story {
	return f.story
}
