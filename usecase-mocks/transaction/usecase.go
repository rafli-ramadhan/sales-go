package transaction

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct{
	mock.Mock
}

func NewTransactionUsecaseMock() *UsecaseMock {
	return &UsecaseMock{}
}

func (m *UsecaseMock) GetTransactionByNumber(number int) (response []model.TransactionDetail, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(number)
	response = ret.Get(0).([]model.TransactionDetail)
	err = ret.Error(1)
	return response, err
}

func (m *UsecaseMock) CreateBulkTransactionDetail(voucherCode string, req model.TransactionDetailBulkRequest) (response []model.TransactionDetail, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(voucherCode, req)
	response = ret.Get(0).([]model.TransactionDetail)
	err = ret.Error(1)
	return response, err
}