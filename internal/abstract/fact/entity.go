package fact

import (
	"github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity interface {
	GetClass() uint
	GetRemark() uint
	GetPosition() uint
	GetStory() uint
	PointToStory(story story.Entity) error
	PointToRemark(remark uint) error
	HasNextFact() (bool, error)
	GetNextFact() (Entity, error)
	GetFirstRemark() (uint, error)
	GetTargetStory() (uint, error)
	ToNextStory(nextFact uint) (Entity, error)
}
