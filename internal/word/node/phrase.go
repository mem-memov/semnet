package node

import "fmt"

type Phrase struct {
	identifier uint
	storage    storage
}

func newPhrase(identifier uint, storage storage) Phrase {
	return Phrase{
		identifier: identifier,
		storage:    storage,
	}
}

func (p Phrase) NewPhrase(word Word) (Phrase, error) {

	identifier, err := p.storage.Create()
	if err != nil {
		return Phrase{}, nil
	}

	err = p.storage.SetReference(word.Identifier(), identifier)
	if err != nil {
		return Phrase{}, nil
	}

	return newPhrase(identifier, p.storage), nil
}

func (p Phrase) Identifier() uint {
	return p.identifier
}

func (p Phrase) Mark(sourceIdentifier uint) error {
	return p.storage.Connect(sourceIdentifier, p.identifier)
}

func (p Phrase) ProvideSingleTarget() (uint, error) {

	targets, err := p.storage.ReadTargets(p.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := p.storage.Create()
		if err != nil {
			return 0, err
		}

		err = p.storage.Connect(p.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", p.identifier, len(targets))
	}
}
