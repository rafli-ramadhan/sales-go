package product

import (
	"errors"
	"sales-go/model"
)

type repository struct {}

func NewRepository() *repository {
	return &repository{}
}

type Repositorier interface {
	GetList() []model.Product
	GetProductByName(name string) (productData model.Product, err error)
	Create(req model.ProductRequest) model.Product
}

func (repo *repository) GetList() []model.Product {
	return model.ProductSlice
}

func (repo *repository) GetProductByName(name string) (productData model.Product, err error) {
	for _, v := range model.ProductSlice {
		if v.Name == name {
			productData = v
		}
	}

	emptyStruct := model.Product{}
	if productData == emptyStruct {
		err = errors.New("Product not found")
		return
	}
	return
}

func (repo *repository) Create(req model.ProductRequest) model.Product {
	newData := model.Product{
		Id:    int64(len(model.ProductSlice)) + 1,
		Name:  req.Name,
		Price: req.Price,
	}
	model.ProductSlice = append(model.ProductSlice, newData)

	return newData
}
