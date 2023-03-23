package model

type Transaction struct {
	Id                int64
	TransactionNumber int64
	Name              string
	Quantity          int64
	Discount          float64
	Total             float64
	Pay				  float64
}

type TransactionDetail struct {
	Id          int64
	Item        string
	Price       float64
	Quantity    int64
	Total       float64
	Transaction Transaction
}

var TransactionSlice []TransactionDetail = []TransactionDetail{}