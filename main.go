package main

import (
	"fmt"
	"sales-go/helpers"

	// handler
	productCtrl "sales-go/handler/product"
	transactionCtrl "sales-go/handler/transaction"
	voucherCtrl "sales-go/handler/voucher"

	// repo
	productRepo "sales-go/repository/product"
	transactionRepo "sales-go/repository/transaction"
	voucherRepo "sales-go/repository/voucher"
)

var isStop bool

func App() {
	// repository
	productRepository := productRepo.NewRepository()
	transactionRepository := transactionRepo.NewRepository()
	voucherRepository := voucherRepo.NewRepository()

	// handler
	productHandler := productCtrl.NewHandler(productRepository)
	transactionHandler := transactionCtrl.NewHandler(transactionRepository, productRepository, voucherRepository)
	voucherHandler := voucherCtrl.NewHandler(voucherRepository)

	var menu int64
	fmt.Println("Choose menu")
	fmt.Println("1. Add New Product")
	fmt.Println("2. Buy a Product")
	fmt.Println("3. Add New Voucher")
	fmt.Println("4. Show List of Product")
	fmt.Println("5. Show List of Voucher")
	fmt.Println("6. Show List of Transaction")
	fmt.Println("7. Get Transaction Detail By Transaction Number")
	fmt.Println("8. Exit")
	fmt.Println("Input menu : ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		productHandler.Create()
		helper.ClearScreeen()
		App()
	case 2:
		transactionHandler.CreateTransaction()
		helper.ClearScreeen()
		App()
	case 3:
		voucherHandler.Create()
		helper.ClearScreeen()
		App()
	case 4:
		productHandler.GetList()
		helper.ClearScreeen()
		App()
	case 5:
		voucherHandler.GetList()
		helper.ClearScreeen()
		App()
	case 6:
		transactionHandler.GetList()
		helper.ClearScreeen()
		App()
	case 7:
		transactionHandler.GetTransactionByNumber()
		helper.ClearScreeen()
		App()
	case 8:
		isStop = true
		return
	}
}

func main() {
	for !isStop {
		App()
	}
}