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
	GetListTransaction() []model.Transaction
	GetTransactionByNumber(transactionNumber int64) (result model.Transaction, err error)
	CreateTransaction(req model.Transaction)
	CreateTransactionDetail(req model.TransactionDetail)
}

func (repo *repository) GetListTransaction() []model.Transaction {
	return model.TransactionSlice
}

func (repo *repository) GetTransactionByNumber(transactionNumber int64) (result model.Transaction, err error) {
	for _, v := range model.TransactionSlice {
		if v.TransactionNumber == transactionNumber {
			return v, nil
		}
	}
	return model.Transaction{}, errors.New("Transaction not found")
}

func (repo *repository) CreateTransaction(req model.Transaction) {
	model.TransactionSlice = append(model.TransactionSlice, req)
}

func (repo *repository) CreateTransactionDetail(req model.TransactionDetail) {
	model.TransactionDetailSlice = append(model.TransactionDetailSlice, req)
}
