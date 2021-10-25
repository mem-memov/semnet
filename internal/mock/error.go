package mock

type ErrorMock struct {
	Error_ func() string
}

func (e ErrorMock) Error() string {
	return e.Error_()
}