package model

type Product struct {
	Id    int64
	Name  string
	Price float64
}

type ProductRequest struct {
	Name  string
	Price float64
}

var ProductSlice []Product = []Product{
	{
		Id:    1,
		Name:  "Kaos_Phincon",
		Price: 150000,
	},
	{
		Id:    2,
		Name:  "Lanyard_Phincon",
		Price: 20000,
	},
	{
		Id:    3,
		Name:  "Tumbler_Phincon",
		Price: 80000,
	},
}