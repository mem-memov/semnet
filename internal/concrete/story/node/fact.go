package node

type Fact struct {
	identifier uint
}

func NewFact(identifier uint) Fact {
	return Fact{
		identifier: identifier,
	}
}

func (f Fact) Identifier() uint {

	return f.identifier
}
