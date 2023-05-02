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
	result = m.Called().Get(0).([]model.TransactionDetail)
	return result, nil
}

func (m *RepoMock) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	res = m.Called().Get(0).([]model.TransactionDetail)
	return res, nil
}
