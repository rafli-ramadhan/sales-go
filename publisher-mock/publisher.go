package publisher

import (
	"github.com/stretchr/testify/mock"
)

type PublisherMock struct {
	mock.Mock
}

func NewPublisher() *PublisherMock {
	return &PublisherMock{}
}

func (m *PublisherMock) Publish(body interface{}) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(body)
	// get return value dari mock
	if ret.Get(0) != nil {
		// type assertion
		err = ret.Get(0).(error)
	}
	return
}
