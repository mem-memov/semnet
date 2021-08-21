package action

import (
	"github.com/mem-memov/semnet/internal/detail"
)

type Action struct {
}

func newAction() Action {
	return Action{}
}

func (a Action) addDetail(name string) (detail.Detail, error) {

}
