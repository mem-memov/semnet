package code

type layer struct {
	storage storage
}

func newLayer(storage storage) *layer {
	return &layer{
		storage: storage,
	}
}

func (l *layer) createEntity(bitNode uint) (Entity, error) {

	codeNode, err := l.storage.GetReference(bitNode)
	if err != nil {
		return Entity{}, err
	}

	if codeNode == 0 {
		codeNode, err = l.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = l.storage.SetReference(bitNode, codeNode)
		if err != nil {
			return Entity{}, err
		}
	}

	characterNode, err := l.storage.GetReference(codeNode)
	if err != nil {
		return Entity{}, err
	}

	if characterNode == 0 {
		characterNode, err = l.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = l.storage.SetReference(codeNode, characterNode)
		if err != nil {
			return Entity{}, err
		}
	}

	return newEntity(bitNode, codeNode, characterNode), nil
}
