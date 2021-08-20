package semnet

import "fmt"

type bit struct {
	node    uint
	storage storage
}

func newBit(node uint, storage storage) bit {
	return bit{
		node:    node,
		storage: storage,
	}
}

func (b bit) getSingleTarget() (uint, error) {

	targets, err := b.storage.ReadTargets(b.node)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := b.storage.Create()
		if err != nil {
			return 0, err
		}

		err = b.storage.Connect(b.node, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("bit %d has too many targets: %d", b.node, len(targets))
	}
}
