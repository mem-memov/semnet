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

	targetIdentifiers, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return []Code{}, err
	}

	if len(targetIdentifiers) > 2 {
		return []Code{}, fmt.Errorf("too many targets in code layer at code %d", c.identifier)
	}

	codes := make([]Code, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		codes[i] = newCode(targetIdentifier, c.storage)
	}

	return codes, nil
}

func (c Code) ReadSources() ([]Code, error) {

	sourceIdentifiers, err := c.storage.ReadSources(c.identifier)
	if err != nil {
		return []Code{}, err
	}

	if len(sourceIdentifiers) > 1 {
		return []Code{}, fmt.Errorf("too many sources in code layer at code %d", c.identifier)
	}

	codes := make([]Code, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		codes[i] = newCode(targetIdentifier, c.storage)
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

	err = c.storage.SetReference(bit.Identifier(), identifier)
	if err != nil {
		return Code{}, err
	}

	err = c.storage.Connect(c.identifier, identifier)
	if err != nil {
		return Code{}, err
	}

	return newCode(identifier, c.storage), nil
}

func (c Code) String() string {
	return fmt.Sprintf("code %d", c.identifier)
}
