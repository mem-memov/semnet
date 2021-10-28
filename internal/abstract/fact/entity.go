package fact

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/remark"
	"github.com/mem-memov/semnet/internal/abstract/story"
)

type Entity interface {
	GetClass() uint
	GetRemark() uint
	GetPosition() uint
	GetStory() uint
	PointToClass(class class.Entity) error
	PointToStory(story story.Entity) error
	PointToRemark(remark remark.Entity) error
	HasNextFact() (bool, error)
	GetNextFact() (Entity, error)
	GetFirstRemark() (uint, error)
}
