package fact

import (
	"github.com/mem-memov/semnet/internal/remark"
)

type Entity struct {

}

func (e Entity) AddRemark(object string, property string) (remark.Entity, error) {

	return remark.Entity{}, nil
}

func (e Entity) GetRemark(remarkIdentifier uint) (remark.Entity, error) {

	return remark.Entity{}, nil
}