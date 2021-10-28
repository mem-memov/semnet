package phrase

import "github.com/mem-memov/semnet/internal/abstract/phrase/node"

type Entity interface {
	GetDetailNode() node.Detail
	HasWordValue(wordValue string) (bool, error)
	WordValue() (string, error)
	FindPrevious(entities Entities) (Entity, bool, error)
	DetailIdentifier() uint
	ProvideDetailTarget(another Entity) (uint, error)
	ProvideNext(wordValue string, entities Entities) (Entity, error)
	Mark(sourceIdentifier uint) error
}
