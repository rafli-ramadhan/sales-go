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
	listProduct = m.Called().Get(0).([]model.Product)
	return listProduct, nil
}

func (m *RepoMock) GetProductByName(name string) (productData model.Product, err error) {
	productData = m.Called().Get(0).(model.Product)
	return productData, nil
}

func (m *RepoMock) Create(req []model.ProductRequest) (result []model.Product, err error) {
	result = m.Called().Get(0).([]model.Product)
	return result, nil
}
