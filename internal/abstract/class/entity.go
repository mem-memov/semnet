package class

type Entity interface {
	CreateBit() (uint, error)
	IsBit(identifier uint) (bool, error)
	GetAllBits() ([]uint, error)

	CreateCharacter() (uint, error)
	IsCharacter(identifier uint) (bool, error)

	CreateWord() (uint, error)
	IsWord(identifier uint) (bool, error)

	CreatePhrase() (uint, error)
	IsPhrase(identifier uint) (bool, error)

	CreateDetail() (uint, error)
	IsDetail(identifier uint) (bool, error)

	CreateRemark() (uint, error)
	IsRemark(identifier uint) (bool, error)

	CreateFact() (uint, error)
	IsFact(identifier uint) (bool, error)

	CreateStory() (uint, error)
	IsStory(identifier uint) (bool, error)
}
