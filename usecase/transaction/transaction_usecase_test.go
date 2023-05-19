package transaction

import (
	"testing"
	"github.com/stretchr/testify/require"

	"sales-go/model"
	productMock "sales-go/mocks/product"
	transactionMock "sales-go/mocks/transaction"
	voucherMock "sales-go/mocks/voucher"
)

func TestGetTransactionByNumber(t *testing.T) {
	t.Run("test get transaction by number", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		transactionNumber := 152309526

		// mock input output
		mockTransactionSuccess.On("GetTransactionByNumber", transactionNumber).Return([]model.TransactionDetail{
			{
				Id: 1,
				Item: "Tumbler_Phincon",
				Price: 30000,
				Quantity: 3,
				Total: 90000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 2,
				Item: "Kaos_Phincon",
				Price: 30000,
				Quantity: 5,
				Total: 150000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 3,
				Item: "Lanyard_Phincon",
				Price: 80000,
				Quantity: 3,
				Total: 240000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
		})

		_, err := usecase.GetTransactionByNumber(transactionNumber)
		require.NoError(t, err)
	})
}

func TestCreateBulkTransactionDetail(t *testing.T) {
	t.Run("test create bulk transaction detail", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1nc0n"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		})
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		})
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		})
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		})
		mockTransactionSuccess.On("CreateBulkTransaction", 152309526, req).Return([]model.TransactionDetail{
			{
				Id: 1,
				Item: "Tumbler_Phincon",
				Price: 30000,
				Quantity: 3,
				Total: 90000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 2,
				Item: "Kaos_Phincon",
				Price: 30000,
				Quantity: 5,
				Total: 150000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 3,
				Item: "Lanyard_Phincon",
				Price: 80000,
				Quantity: 3,
				Total: 240000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
		})

		_, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.NoError(t, err)
	})
}