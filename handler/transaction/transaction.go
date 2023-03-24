package transaction

import (
	"fmt"
	"math/rand"
	"sales-go/model"
	"sales-go/repository/product"
	"sales-go/repository/transaction"
	"sales-go/repository/voucher"
	"time"
)

type handler struct {
	repo        transaction.Repositorier
	productrepo product.Repositorier
	voucherrepo voucher.Repositorier
}

func NewHandler(
	repositorier transaction.Repositorier,
	productRepository product.Repositorier,
	voucherRepository voucher.Repositorier,
) *handler {
	return &handler{
		repo:        repositorier,
		productrepo: productRepository,
		voucherrepo: voucherRepository,
	}
}

type Handlerer interface {
	GetList()
	GetTransactionByNumber()
	CreateTransaction()
}

func (handler *handler) GetList() {
	result := handler.repo.GetListTransaction()
	fmt.Printf("\nId\t|TransactionNumber\t|Name\t\t|Quantity\t\t|Discount\t\t|Total\t\t|Pay")
	for _, v := range result {
		fmt.Printf("\n%d\t|%d\t\t|%s\t\t|%d\t\t|%0.2f\t\t|%f\t\t|%0.2f", v.Transaction.Id, v.Transaction.TransactionNumber, v.Transaction.Name, v.Transaction.Quantity, v.Transaction.Discount, v.Transaction.Total, v.Transaction.Pay)
	}
}

func (handler *handler) GetTransactionByNumber() {
	var transactionNumber int64
	fmt.Println("\nInput transaction number : ")
	fmt.Scanln(&transactionNumber)
	
	result, err := handler.repo.GetTransactionByNumber(transactionNumber)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println("\nTOKO PHINCON")
	fmt.Println("Jl. Arteri Pd. Indah - Jakarta")
	fmt.Printf("Transaction Number %d\n", result.Transaction.TransactionNumber)
	fmt.Println("--------------------------------------\n")
	fmt.Printf("%s\t\t\n", result.Transaction.Name)
	fmt.Printf("Rp.%0.2f\t\tx%d\n", result.Price, result.Transaction.Quantity)
	fmt.Printf("\nDiscount\t\t%0.2f persen\n", result.Transaction.Discount)
	fmt.Printf("Total\t\t\tRp.%0.0f\n", result.Transaction.Total)
	fmt.Printf("Pay\t\t\tRp.%0.2f\n", result.Transaction.Pay)
	fmt.Printf("Revenue\t\t\tRp.%0.2f\n", result.Transaction.Pay-result.Transaction.Total)
}

func (handler *handler) CreateTransaction() {
	var name string
	var quantity int64
	var total float64
	var pay float64

	// 1. input product name
	fmt.Println("\nInput product name : ")
	fmt.Scanln(&name)

	// 2. search product
	product, err := handler.productrepo.GetProductByName(name)
	if err != nil {	
		fmt.Println("\nSorry, the product you are looking for does not exist.")
		fmt.Println("Here are the list of products")
		handler.productrepo.GetList()

		handler.CreateTransaction()
	}

	// 3. input quantity
	fmt.Println("\nInput quantity : ")
	fmt.Scanln(&quantity)

	if quantity <= 0 {
		fmt.Println("Product quantity should be positive number and not 0.")

		handler.CreateTransaction()
	}

	total = float64(quantity)*product.Price

	var discount float64 = 0
	if total > 300000 {
		var voucherCode string
		fmt.Println("\nInput voucher code : ")
		fmt.Scanln(&voucherCode)

		voucher, err := handler.voucherrepo.GetVoucherByCode(voucherCode)
		if err != nil {
			fmt.Println(err)

			fmt.Println("\nSorry, there is no voucher with name %s\n", voucherCode)
		} else {
			discount = voucher.Persen/100
			total = total*(discount)

			fmt.Println("\nCongratulation, there is a discount ", voucher.Persen)
		}
	}

	// 4. Show total user should pay
	fmt.Printf("\nTotal price you should pay : %0.2f\n", total)

	// 5. input pay
	fmt.Println("\nInput the nominal you want to pay : ")
	fmt.Scanln(&pay)

	if pay <= 0 {
		fmt.Println("Input pay should be positive number and not 0.")

		handler.CreateTransaction()
	}
	
	// 6. calculate refund
	fmt.Println("\nRefund : ", pay - total)

	// 7. Input new transaction to transaction detail
	newTransaction := model.Transaction{
		Id:                int64(len(model.TransactionSlice)) + 1,
		TransactionNumber: int64(rand.Intn(10000000000)),
		Name:              name,
		Quantity:          quantity,
		Discount:          discount,
		Total:             total,
		Pay:               pay,
	}

	// 1. Buat transaction detail berulang-ulang tanpa Transaction, 2. Buat transaction, 3. Transaction detail yang sudah dibuat, dimasukkan struct transactionnya.
	newTransactionDetail := model.TransactionDetail{
		Id:          int64(len(model.TransactionSlice)) + 1,
		Item:        name,
		Price:       product.Price,
		Quantity:    quantity,
		Total:       total,
		Transaction: newTransaction,
	}

	// 8. Input transaction detail to transaction slice
	result := handler.repo.CreateTransactionDetail(newTransactionDetail)

	// 9. Show transaction struct
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	formatTime := now.In(loc).Format("02/01/2006 15:04")

	fmt.Println("\n==== Transaction Success ===\n")
	fmt.Println("\nTOKO PHINCON")
	fmt.Println("Jl. Arteri Pd. Indah - Jakarta")
	fmt.Printf("%s\n", formatTime)
	fmt.Printf("Transaction Number %d\n", result.Transaction.TransactionNumber)
	fmt.Println("--------------------------------------\n")
	fmt.Printf("%s\t\t\n", result.Transaction.Name)
	fmt.Printf("Rp.%0.2f\t\tx%d\n", result.Price, result.Transaction.Quantity)
	fmt.Printf("\nDiscount\t\t%0.2f persen\n", result.Transaction.Discount)
	fmt.Printf("Total\t\t\tRp.%0.0f\n", result.Transaction.Total)
	fmt.Printf("Pay\t\t\tRp.%0.2f\n", result.Transaction.Pay)
	fmt.Printf("Revenue\t\t\tRp.%0.2f\n", result.Transaction.Pay-result.Transaction.Total)
}