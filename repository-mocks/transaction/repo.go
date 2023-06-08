package transaction

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct{
	mock.Mock
}

func NewTransactionRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(transactionNumber)
	// get return value dari mock
	result = ret.Get(0).([]model.TransactionDetail)
	err = ret.Error(1)
	return
}

func (m *RepoMock) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {// sebagai indikator parameter diperoleh
	// sebagai indikator parameter diperoleh
	ret := m.Called(voucher, listTransactionDetail, req)
	// get return value dari mock
	res = ret.Get(0).([]model.TransactionDetail)
	err = ret.Error(1)
	return
}
