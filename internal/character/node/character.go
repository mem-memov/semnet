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

func (c Character) ReadTargets() ([]Character, error) {

	targetIdentifiers, err := c.storage.ReadTargets(c.identifier)
	if err != nil {
		return []Character{}, err
	}

	if len(targetIdentifiers) > 2 {
		return []Character{}, fmt.Errorf("too many targets in character layer at character %d", c.identifier)
	}

	characters := make([]Character, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		characters[i] = newCharacter(targetIdentifier, c.storage)
	}

	return characters, nil
}

func (c Character) ReadSources() ([]Character, error) {

	sourceIdentifiers, err := c.storage.ReadSources(c.identifier)
	if err != nil {
		return []Character{}, err
	}

	if len(sourceIdentifiers) > 1 {
		return []Character{}, fmt.Errorf("too many sources in character layer at character %d", c.identifier)
	}

	characters := make([]Character, len(sourceIdentifiers))

	for i, targetIdentifier := range sourceIdentifiers {
		characters[i] = newCharacter(targetIdentifier, c.storage)
	}

	return characters, nil
}

func (c Character) GetClassAndBitAndWord() (uint, uint, uint, error) {

	bitIdentifier, wordIdentifier, err := c.storage.GetReference(c.identifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	classIdentifier, characterIdentifier, err := c.storage.GetReference(bitIdentifier)
	if err != nil {
		return 0, 0, 0, nil
	}

	if characterIdentifier != c.identifier {
		return 0, 0, 0, fmt.Errorf("cheracter entity invalid at charcter node %d", c.identifier)
	}

	return classIdentifier, bitIdentifier, wordIdentifier, nil
}

func (c Character) NewCharacter(bit Bit) (Character, error) {

	identifier, err := c.storage.Create()
	if err != nil {
		return Character{}, err
	}

	err = c.storage.SetReference(bit.Identifier(), identifier)
	if err != nil {
		return Character{}, err
	}

	err = c.storage.Connect(c.identifier, identifier)
	if err != nil {
		return Character{}, err
	}

	return newCharacter(identifier, c.storage), nil
}

func (c Character) String() string {
	return fmt.Sprintf("character %d", c.identifier)
}
