package product

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct{
	mock.Mock
}

func NewProductRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) GetList() (listProduct []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	// get return value dari mock
	listProduct = ret.Get(0).([]model.Product)
	err = ret.Error(1)
	return
}

func (m *RepoMock) GetProductByName(name string) (productData model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	// get return value dari mock
	productData = ret.Get(0).(model.Product)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *RepoMock) Create(req []model.ProductRequest) (result []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	// get return value dari mock
	result = ret.Get(0).([]model.Product)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}
