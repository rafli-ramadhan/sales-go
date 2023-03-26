package model

type Transaction struct {
	Id                int64
	TransactionNumber int64
	Quantity          int64
	Discount          float64
	Total             float64
	Pay				  float64
	TransactionDetail []TransactionDetail
}

type TransactionDetail struct {
	Id          int64
	Item        string
	Price       float64
	Quantity    int64
	Total       float64
}

var Total float64

var TransactionSlice []Transaction = []Transaction{}

var TransactionDetailSlice []TransactionDetail = []TransactionDetail{}
