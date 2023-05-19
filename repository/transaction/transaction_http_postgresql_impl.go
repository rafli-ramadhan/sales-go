package transaction

import (
	"context"
	"time"

	"sales-go/db"
	"sales-go/model"
	"sales-go/publisher"
)

type repositoryhttppostgresql struct {}

func NewPostgreSQLHTTPRepository() *repositoryhttppostgresql {
	return &repositoryhttppostgresql{}
}

func (repo *repositoryhttppostgresql) GetTransactionByNumber(transactionNumber int) (result []model.TransactionDetail, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// GetTransaction
	query := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, transactionNumber)
	if err != nil {
		return
	}

	var transaction model.Transaction
	for res.Next() {
		res.Scan(&transaction.Id, &transaction.TransactionNumber, &transaction.Name, &transaction.Quantity, &transaction.Discount, &transaction.Total, &transaction.Pay)
	}

	// GetTransactionDetail
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`
	stmt2, err := db.PrepareContext(ctx, query2)
	if err != nil {
		return
	}

	res2, err := stmt2.QueryContext(ctx, transaction.Id)
	if err != nil {
		return
	}

	for res2.Next() {
		var temp model.TransactionDetail
		res2.Scan(&temp.Id, &temp.Item, &temp.Price, &temp.Quantity, &temp.Total)
		// append transaction in each of transaction detail
		temp.Transaction = transaction
		result = append(result, temp)
	}

	return
}

type TransactionData struct {
	Voucher 				model.VoucherRequest
	ListTransactionDetail 	[]model.TransactionDetail
	Req 					model.TransactionDetailBulkRequest
}

func (repo *repositoryhttppostgresql) CreateBulkTransactionDetail(voucher model.VoucherRequest, listTransactionDetail []model.TransactionDetail, req model.TransactionDetailBulkRequest) (res []model.TransactionDetail, err error) {
	sendingData := TransactionData{
		voucher, 
		listTransactionDetail, 
		req,
	}

	err = publisher.Publish(sendingData)
	return
}

