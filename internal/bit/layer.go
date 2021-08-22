package bit

import "fmt"

type layer struct {
	storage       storage
	isInitialized bool
}

func newLayer(storage storage) *layer {
	return &layer{
		storage:       storage,
		isInitialized: false,
	}
}

func (l *layer) initialize() error {

	if l.isInitialized {
		return nil
	}

	hasZero, err := l.storage.Has(bitZeroNode)
	if err != nil {
		return err
	}

	if !hasZero {
		zeroNode, err := l.storage.Create()
		if err != nil {
			return err
		}

		if zeroNode != bitZeroNode {
			return fmt.Errorf("invalid zero identifier %d", zeroNode)
		}
	}

	hasOne, err := l.storage.Has(bitOneNode)
	if err != nil {
		return err
	}

	if !hasOne {
		oneNode, err := l.storage.Create()
		if err != nil {
			return err
		}

		if oneNode != bitOneNode {
			return fmt.Errorf("invalid one identifier %d", oneNode)
		}
	}

	l.isInitialized = true

	return nil
}
