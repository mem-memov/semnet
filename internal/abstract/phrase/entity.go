package phrase

import "github.com/mem-memov/semnet/internal/abstract/phrase/node"

type Entity interface {
	GetClass() uint
	GetWord() uint
	GetPhrase() uint
	GetDetail() uint
	PointToPhrase(phrase uint) error
	GetTargetPhrases() ([]Entity, error)
	GetTargetWord() (uint, error)

	GetDetailNode() node.Detail
	HasWordValue(wordValue string) (bool, error)
	WordValue() (string, error)
	FindPrevious(entities Entities) (Entity, bool, error)
	DetailIdentifier() uint
	ProvideDetailTarget(another Entity) (uint, error)
	Mark(sourceIdentifier uint) error
}
