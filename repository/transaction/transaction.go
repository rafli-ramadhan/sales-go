package transaction

import (
	"errors"
	"sales-go/model"
)

type repository struct {}

func NewRepository() *repository {
	return &repository{}
}

type Repositorier interface {
	GetListTransaction() []model.TransactionDetail
	GetTransactionByNumber(transactionNumber int64) (result model.TransactionDetail, err error)
	CreateTransactionDetail(req model.TransactionDetail) (result model.TransactionDetail)
}

func (repo *repository) GetListTransaction() []model.TransactionDetail {
	return model.TransactionSlice
}

func (repo *repository) GetTransactionByNumber(transactionNumber int64) (result model.TransactionDetail, err error) {
	for _, v := range model.TransactionSlice {
		if v.Transaction.TransactionNumber == transactionNumber {
			return v, nil
		}
	}
	return model.TransactionDetail{}, errors.New("Transaction not found")
}

func (repo *repository) CreateTransactionDetail(req model.TransactionDetail) (result model.TransactionDetail) {
	model.TransactionSlice = append(model.TransactionSlice, req)
	result = req
	return
}
