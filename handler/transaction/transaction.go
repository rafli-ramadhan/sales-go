package transaction

import (
	"fmt"
	"math/rand"
	"sales-go/model"
	"sales-go/repository/product"
	"sales-go/repository/transaction"
	"sales-go/repository/voucher"
	"strings"
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
	fmt.Printf("\nId\t|TransactionNumber")
	for _, v := range result {
		fmt.Printf("\n%d\t|%d\t\t", v.Id, v.TransactionNumber)
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
	fmt.Println("--------------------------------------\n")
	fmt.Printf("Transaction Number\t%d\n", result.TransactionNumber)
	for _, v := range result.TransactionDetail {
		fmt.Printf("\n%s\n", v.Item)
		fmt.Printf("Rp.%0.2f\t\tx%d\n", v.Price, v.Quantity)
	}
	fmt.Printf("Qauantity\t\t%d\n", result.Quantity)
	fmt.Printf("Discount\t\t%0.2f persen\n", result.Discount)
	fmt.Printf("Total\t\t\tRp.%0.0f\n", result.Total)
	fmt.Printf("Pay\t\t\tRp.%0.2f\n", result.Pay)
	fmt.Printf("Revenue\t\t\tRp.%0.2f\n", result.Pay-result.Total)
}

func (handler *handler) CreateTransaction() {
	// input product name
	var name string
	var product model.Product
	var err error
	for {
		fmt.Println("\n\nInput product name : ")
		fmt.Scanln(&name)

		// search product
		product, err = handler.productrepo.GetProductByName(name)
		if err != nil {	
			fmt.Println("\nSorry, the product you are looking for does not exist.")
			fmt.Println("\nHere are the list of products")
			result := handler.productrepo.GetList()
			fmt.Printf("\nId\t\t|Name\t\t\t|Price\t\t")
			for _, v := range result {
				if len(v.Name) > 13 {
					fmt.Printf("\n%d\t\t|%s\t|%0.2f", v.Id, v.Name, v.Price)
				} else if len(v.Name) > 5 && len(v.Name) < 13 {
					fmt.Printf("\n%d\t\t|%s\t\t|%0.2f", v.Id, v.Name, v.Price)
				} else {
					fmt.Printf("\n%d\t\t|%s\t\t\t|%0.2f", v.Id, v.Name, v.Price)
				}
			}
		} else {
			break
		}
	}

	// input quantity
	var quantity int64
	for {
		fmt.Println("\nInput quantity : ")
		fmt.Scanln(&quantity)
		if quantity <= 0 {
			fmt.Println("Product quantity should be positive number and not 0.")
		} else {
			break
		}
	}

	newTransactionDetail := model.TransactionDetail{
		Id:          int64(len(model.TransactionSlice)) + 1,
		Item:        name,
		Price:       product.Price,
		Quantity:    quantity,
		Total:       float64(quantity)*product.Price,
	}
	// input transaction detail to transaction slice
	handler.repo.CreateTransactionDetail(newTransactionDetail)

	model.Total = model.Total + float64(quantity)*product.Price
	
	var response string
	fmt.Println("\nDo you want to buy another product ? (y | n)")
	fmt.Scanln(&response)
	if strings.ToLower(response) != "n" && strings.ToLower(response) != "no" {
		handler.CreateTransaction()
	}

	// calculate total
	var discount float64
	var voucherPersen float64
	if model.Total > 300000 {
		for {
			var response string
			fmt.Println("\nDo you have voucher code ? (y | n)")
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
				break
			} else {
				for {
					var voucherCode string
					fmt.Println("\nInput voucher code : ")
					fmt.Scanln(&voucherCode)
		
					voucher, err := handler.voucherrepo.GetVoucherByCode(voucherCode)
					if err != nil {
						fmt.Println(err)
						fmt.Printf("\nSorry, there is no voucher with name %s.\n", voucherCode)
					} else {
						voucherPersen = voucher.Persen
						discount = voucherPersen/100
						model.Total = model.Total*(1-discount)
						fmt.Printf("\nCongratulation, there is a discount %0.2f persen.\n", voucher.Persen)
						break
					}
				}
				break
			}
		}
	}

	// show total user should pay
	fmt.Printf("\nTotal price you should pay : %0.2f\n", model.Total)

	// input pay
	var pay float64
	for {
		fmt.Println("\nInput the nominal you want to pay : ")
		fmt.Scanln(&pay)
		if pay <= 0 {
			fmt.Println("Input pay should be positive number and not 0.")
		} else if pay < model.Total {
			fmt.Println("Not enough money")
		} else {
			break
		}
	}
	
	// calculate refund
	fmt.Println("\nRefund : ", pay - model.Total)

	// input new transaction to transaction detail
	newTransaction := model.Transaction{
		Id:                int64(len(model.TransactionSlice)) + 1,
		TransactionNumber: int64(rand.Intn(10000000000)),
		Quantity:          quantity,
		Discount:          voucherPersen,
		Total:             model.Total,
		Pay:               pay,
		TransactionDetail: model.TransactionDetailSlice,
	}
	// input transaction detail to transaction slice
	handler.repo.CreateTransaction(newTransaction)

	// re-empty total
	model.Total = 0

	// show transaction struct
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
	fmt.Println("--------------------------------------")
	fmt.Printf("Transaction Number %d\n", newTransaction.TransactionNumber)
	for _, v := range newTransaction.TransactionDetail {
		fmt.Printf("\n%s\n", v.Item)
		fmt.Printf("Rp.%0.2f\t\tx%d\n", v.Price, v.Quantity)
	}
	fmt.Printf("\nDiscount\t\t%0.2f persen\n", newTransaction.Discount)
	fmt.Printf("Total\t\t\tRp.%0.2f\n", newTransaction.Total)
	fmt.Printf("Pay\t\t\tRp.%0.2f\n", newTransaction.Pay)
	fmt.Printf("Revenue\t\t\tRp.%0.2f\n", newTransaction.Pay - newTransaction.Total)
}