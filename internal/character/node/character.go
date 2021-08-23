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
		return []Character{}, fmt.Errorf("too many targets in ccharacter layer at ccharacter %d", c.identifier)
	}

	characters := make([]Character, len(targetIdentifiers))

	for i, targetIdentifier := range targetIdentifiers {
		characters[i] = newCharacter(targetIdentifier, c.storage)
	}

	return characters, nil
}
