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
	response = ret.Get(0).([]model.Product)
	err = ret.Error(1)
	return response, err
}

func (m *UsecaseMock) GetProductByName(name string) (response model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	response = ret.Get(0).(model.Product)
	err = ret.Error(1)
	return response, err
}

func (m *UsecaseMock) Create(req []model.ProductRequest) (response []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	response = ret.Get(0).([]model.Product)
	err = ret.Error(1)
	return response, err
}