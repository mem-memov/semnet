package class

import abstractClass "github.com/mem-memov/semnet/internal/abstract/class"

type EntityMock struct {
	CreateBit_  func() (uint, error)
	IsBit_      func(identifier uint) (bool, error)
	GetAllBits_ func() ([]uint, error)

	CreateCharacter_ func() (uint, error)
	IsCharacter_     func(identifier uint) (bool, error)

	CreateWord_ func() (uint, error)
	IsWord_     func(identifier uint) (bool, error)

	CreatePhrase_ func() (uint, error)
	IsPhrase_     func(identifier uint) (bool, error)

	CreateDetail_ func() (uint, error)
	IsDetail_     func(identifier uint) (bool, error)

	CreateRemark_ func() (uint, error)
	IsRemark_     func(uint) (bool, error)

	CreateFact_ func() (uint, error)
	IsFact_     func(identifier uint) (bool, error)

	CreateStory_ func() (uint, error)
	IsStory_     func(identifier uint) (bool, error)
}

var _ abstractClass.Entity = EntityMock{}

func (e EntityMock) CreateBit() (uint, error) {
	return e.CreateBit_()
}

func (e EntityMock) IsBit(identifier uint) (bool, error) {
	return e.IsBit_(identifier)
}

func (e EntityMock) GetAllBits() ([]uint, error) {
	return e.GetAllBits()
}

func (e EntityMock) CreateCharacter() (uint, error) {
	return e.CreateCharacter_()
}

func (e EntityMock) IsCharacter(identifier uint) (bool, error) {
	return e.IsCharacter_(identifier)
}

func (e EntityMock) CreateWord() (uint, error) {
	return e.CreateWord_()
}

func (e EntityMock) IsWord(identifier uint) (bool, error) {
	return e.IsWord_(identifier)
}

func (e EntityMock) CreatePhrase() (uint, error) {
	return e.CreatePhrase_()
}

func (e EntityMock) IsPhrase(identifier uint) (bool, error) {
	return e.IsPhrase_(identifier)
}

func (e EntityMock) CreateDetail() (uint, error) {
	return e.CreateDetail_()
}

func (e EntityMock) IsDetail(identifier uint) (bool, error) {
	return e.IsDetail_(identifier)
}

func (e EntityMock) CreateRemark() (uint, error) {
	return e.CreateRemark_()
}

func (e EntityMock) IsRemark(identifier uint) (bool, error) {
	return e.IsRemark_(identifier)
}

func (e EntityMock) CreateFact() (uint, error) {
	return e.CreateFact_()
}

func (e EntityMock) IsFact(identifier uint) (bool, error) {
	return e.IsFact_(identifier)
}

func (e EntityMock) CreateStory() (uint, error) {
	return e.CreateStory_()
}

func (e EntityMock) IsStory(identifier uint) (bool, error) {
	return e.IsStory_(identifier)
}
