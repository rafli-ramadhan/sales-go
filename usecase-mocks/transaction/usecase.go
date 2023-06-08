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
	// get return value dari mock
	response = ret.Get(0).([]model.TransactionDetail)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return response, err
}

func (m *UsecaseMock) CreateBulkTransactionDetail(voucherCode string, req model.TransactionDetailBulkRequest) (response []model.TransactionDetail, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(voucherCode, req)
	// get return value dari mock
	response = ret.Get(0).([]model.TransactionDetail)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return response, err
}