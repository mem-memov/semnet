package node

import "fmt"

type Code struct {
	identifier uint
	storage    storage
}

func newCode(identifier uint, storage storage) Code {
	return Code{
		identifier: identifier,
		storage:    storage,
	}
}

func (c Code) Identifier() uint {
	return c.identifier
}

func (c Code) ReadTargets() ([]Code, error) {

	targets, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return []Code{}, err
	}

	if len(targets) > 2 {
		return []Code{}, fmt.Errorf("too many targets in code layer at code %d", c.identifier)
	}

	codes := make([]Code, len(targets))

	for i, target := range targets {
		codes[i] = newCode(target, c.storage)
	}

	return codes, nil
}

func (c Code) GetBitAndCharacter() (uint, uint, error) {

	bitIdentifier, characterIdentifier, err := c.storage.GetReference(c.identifier)
	if err != nil {
		return 0, 0, nil
	}

	return bitIdentifier, characterIdentifier, nil
}

func (c Code) NewCode(bit Bit) (Code, error) {

	identifier, err := c.storage.Create()
	if err != nil {
		return Code{}, err
	}

	err = c.storage.SetReference(identifier, bit.Identifier())
	if err != nil {
		return Code{}, err
	}

	return newCode(identifier, c.storage), nil
}
