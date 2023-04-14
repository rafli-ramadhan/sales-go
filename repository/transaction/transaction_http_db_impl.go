package transaction

import (
	"sales-go/model"
)

type repositoryhttpdb struct {}

func NewDBHTTPRepository() *repositoryhttpdb {
	return &repositoryhttpdb{}
}

func (repo *repositoryhttpdb) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	// query := `SELECT FROM transaction WHERE `
	return
}

func (repo *repositoryhttpdb) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	// query := `INSERT INTO transaction () values (?, ?, ?, ?)`
	return
}