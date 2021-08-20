package semnet

type Graph struct {
	bits    *bits
	actions *actions
}

func NewGraph(storage storage) *Graph {
	bits := newBits(storage)
	codes := newCodes(storage, bits)
	characters := newCharacters(storage, codes)
	words := newWords(storage, characters)
	actions := newActions(storage, words)

	return &Graph{
		bits:    bits,
		actions: actions,
	}
}

func (g *Graph) InitializeBits() error {
	return g.bits.initialize()
}

func (g *Graph) addAction(name string) (Action, error) {
	return g.actions.create(name)
}
