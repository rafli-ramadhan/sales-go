package product

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct{
	mock.Mock
}

func NewProductUsecaseMock() *UsecaseMock {
	return &UsecaseMock{}
}

func (m *UsecaseMock) GetList() (response []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	// get return value dari mock
	response = ret.Get(0).([]model.Product)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *UsecaseMock) GetProductByName(name string) (response model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	// get return value dari mock
	response = ret.Get(0).(model.Product)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *UsecaseMock) Create(req []model.ProductRequest) (response []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	// get return value dari mock
	response = ret.Get(0).([]model.Product)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}