package node

import "fmt"

type Character struct {
	identifier uint
	storage    storage
}

func newCharacter(identifier uint, storage storage) Character {
	return Character{
		identifier: identifier,
		storage:    storage,
	}
}

func (c Character) Identifier() uint {
	return c.identifier
}

func (c Character) mark(sourceIdentifier uint) error {
	return c.storage.Connect(sourceIdentifier, c.identifier)
}

func (c Character) ProvideSingleTarget() (uint, error) {

	targets, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := c.storage.Create()
		if err != nil {
			return 0, err
		}

		err = c.storage.Connect(c.identifier, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("entity %d has too many targets: %d", c.identifier, len(targets))
	}
}

func (c Character) NewCharacter(code Code) (Character, error) {

	identifier, err := c.storage.Create()
	if err != nil {
		return Character{}, nil
	}

	err = c.storage.SetReference(code.Identifier(), identifier)
	if err != nil {
		return Character{}, nil
	}

	return newCharacter(identifier, c.storage), nil
}

func (c Character) String() string {
	return fmt.Sprintf("character %d", c.identifier)
}
