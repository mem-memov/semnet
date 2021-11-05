package phrase

import (
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type EntityMock struct {
	GetClass_         func() uint
	GetWord_          func() uint
	GetPhrase_        func() uint
	GetDetail_        func() uint
	PointToPhrase_    func(uint) error
	GetTargetPhrases_ func() ([]abstractPhrase.Entity, error)
	GetSourceWord_    func() (uint, error)
	HasSourcePhrase_  func() (bool, error)
	GetSourcePhrase_  func() (abstractPhrase.Entity, error)
	GetSourceDetails_ func() ([]uint, error)
	GetTargetDetails_ func() ([]uint, error)
	AddSourceDetail_  func(detail uint) error
	AddTargetDetail_  func(detail uint) error
	Mark_             func(sourceIdentifier uint) error
}

var _ abstractPhrase.Entity = EntityMock{}

func (e EntityMock) GetClass() uint {

	return e.GetClass_()
}

func (e EntityMock) GetWord() uint {

	return e.GetWord_()
}

func (e EntityMock) GetPhrase() uint {

	return e.GetPhrase_()
}

func (e EntityMock) GetDetail() uint {

	return e.GetDetail_()
}

func (e EntityMock) PointToPhrase(phrase uint) error {

	return e.PointToPhrase_(phrase)
}

func (e EntityMock) GetTargetPhrases() ([]abstractPhrase.Entity, error) {

	return e.GetTargetPhrases_()
}

func (e EntityMock) GetSourceWord() (uint, error) {

	return e.GetSourceWord_()
}

func (e EntityMock) HasSourcePhrase() (bool, error) {

	return e.HasSourcePhrase_()
}

func (e EntityMock) GetSourcePhrase() (abstractPhrase.Entity, error) {

	return e.GetSourcePhrase_()
}

func (e EntityMock) GetSourceDetails() ([]uint, error) {

	return e.GetSourceDetails_()
}

func (e EntityMock) GetTargetDetails() ([]uint, error) {

	return e.GetTargetDetails_()
}

func (e EntityMock) AddSourceDetail(detail uint) error {

	return e.AddSourceDetail_(detail)
}

func (e EntityMock) AddTargetDetail(detail uint) error {

	return e.AddTargetDetail_(detail)
}

func (e EntityMock) Mark(sourceIdentifier uint) error {

	return e.Mark_(sourceIdentifier)
}
