package bit

import (
	"fmt"
)

type Entity struct {
	node    uint
	storage storage
}

func newEntity(node uint, storage storage) Entity {
	return Entity{
		node:    node,
		storage: storage,
	}
}

func (e Entity) Identifier() uint {
	return e.node
}

func (e Entity) Is(bit bool) bool {
	if bit {
		return e.node == bitOneNode
	} else {
		return e.node == bitZeroNode
	}
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	targets, err := e.storage.ReadTargets(e.node)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := e.storage.Create()
		if err != nil {
			return 0, err
		}

		err = e.storage.Connect(e.node, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", e.node, len(targets))
	}
}
