package product

import (
	"fmt"
	"sales-go/model"
	"sales-go/repository/product"
)

type handler struct {
	repo product.Repositorier
}

func NewHandler(repositorier product.Repositorier) *handler {
	return &handler{
		repo: repositorier,
	}
}

type handlerer interface {
	GetList()
	Create()
}

func (handler *handler) GetList() {
	result := handler.repo.GetList()
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
}

func (handler *handler) Create() {
	var name string
	var price float64
	fmt.Println("\nInput name data : ")
	fmt.Scanln(&name)
	fmt.Println("\nInput price data : ")
	fmt.Scanln(&price)

	if price <= 0 {
		fmt.Println("Product price should be positive number and not 0.")

		handler.Create()
	}

	result := handler.repo.Create(model.ProductRequest{
		Name:  name,
		Price: price,
	})

	fmt.Println("Product has been added with id : ", result.Id)
}