package product

import (
	"testing"
	"github.com/stretchr/testify/require"

	"sales-go/mocks/product"
	"sales-go/model"
)

func TestGetList(t *testing.T) {
	t.Run("test get list product", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		mockSuccess.On("GetList").Return([]model.Product{
			{
				Id: 7,
				Name: "Kaos_Phincon_2",
				Price: 30000,
			},
			{
				Id: 8,
				Name: "Lanyard_Phincon_2",
				Price: 80000,
			},
			{
				Id: 9,
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		})

		_, err := usecase.GetList()
		require.NoError(t, err)
	})
}

func TestGetProductByName(t *testing.T) {
	t.Run("test get product by name", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)
		
		productName := "Kaos Phincon"
		mockSuccess.On("GetProductByName", productName).Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		})

		_, err := usecase.GetProductByName(productName)
		require.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("test create product", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon_2",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon_2",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}
		mockSuccess.On("Create", req).Return([]model.Product{
			{
				Id: 7,
				Name: "Kaos_Phincon_2",
				Price: 30000,
			},
			{
				Id: 8,
				Name: "Lanyard_Phincon_2",
				Price: 80000,
			},
			{
				Id: 9,
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		})

		_, err := usecase.Create(req)
		require.NoError(t, err)
	})
}
