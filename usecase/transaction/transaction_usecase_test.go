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
		
		productName := "Kaos Phincon"
		mockProductSuccess.On("GetProductByName", productName).Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		})

		transactionNumber := 152309526
		mockTransactionSuccess.On("GetTransactionByNumber", transactionNumber).Return(model.TransactionDetail{
		})

		code := "Ph1ncon"
		mockVoucherSuccess.On("GetVoucherByCode", code).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		})
		
		_, err := usecase.GetTransactionByNumber(transactionNumber)
		require.NoError(t, err)
	})
}
