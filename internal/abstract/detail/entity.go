package detail

import "github.com/mem-memov/semnet/internal/abstract/remark"

type Entity interface {
	HasPhraseValue(phraseValue string) (bool, error)
	PointToRemark(remark remark.Entity) error
	GetObjectAndProperty() (string, string, error)
}
