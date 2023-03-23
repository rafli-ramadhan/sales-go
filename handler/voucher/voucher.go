package voucher

import (
	"fmt"
	"sales-go/model"
	"sales-go/repository/voucher"
)

type handler struct {
	repo voucher.Repositorier
}

func NewHandler(repositorier voucher.Repositorier) *handler {
	return &handler{
		repo: repositorier,
	}
}

type Handlerer interface {
	GetList()
	Create()
}

func (handler *handler) GetList() {
	result := handler.repo.GetList()
	fmt.Printf("\nId\t\t|Code\t\t\t|Persen\t\t")
	for _, v := range result {
		if len(v.Code) > 13 {
			fmt.Printf("\n%d\t\t|%s\t|%0.2f", v.Id, v.Code, v.Persen)
		} else if len(v.Code) > 5 && len(v.Code) < 13 {
			fmt.Printf("\n%d\t\t|%s\t\t|%0.2f", v.Id, v.Code, v.Persen)
		} else {
			fmt.Printf("\n%d\t\t|%s\t\t\t|%0.2f", v.Id, v.Code, v.Persen)
		}
	}
}

func (handler *handler) Create() {	
	var code string
	var persen float64
	fmt.Println("\nInput code data : ")
	fmt.Scanln(&code)
	fmt.Println("\nInput persen data : ")
	fmt.Scanln(&persen)

	handler.repo.Create(model.VoucherRequest{
		Code:   code,
		Persen: persen,
	})
}