package transaction

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"sales-go/model"
	"sales-go/usecase-mocks/transaction"
)

func TestHandlerGin(t *testing.T) {
	t.Run("get transaction by number success", func(t *testing.T) {
		ucMock := transaction.NewTransactionUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetTransactionByNumber", 332980303).Return(
			[]model.TransactionDetail{
				{
					Id: 43,
					Item: "Tumbler_Phincon",
					Price: 30000,
					Quantity: 3,
					Total: 90000,
					Transaction: model.Transaction{
						Id: 24,
						TransactionNumber: 332980303,
						Name: "Utsman",
						Quantity: 11,
						Discount: 0.3,
						Total: 336000,
						Pay: 1000000,
					},
				},
				{
					Id: 44,
					Item: "Kaos_Phincon_2",
					Price: 30000,
					Quantity: 5,
					Total: 150000,
					Transaction: model.Transaction{
						Id: 24,
						TransactionNumber: 332980303,
						Name: "Utsman",
						Quantity: 11,
						Discount: 0.3,
						Total: 336000,
						Pay: 1000000,
					},
				},
			}, nil)


		transactionNumber := "332980303"
		URL := fmt.Sprintf("http://localhost:5000/transaction?transaction_id=%s", transactionNumber)
		request := httptest.NewRequest(http.MethodGet, URL, nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetTransactionByNumber(c)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("get transaction by number failed : null query param transaction id", func(t *testing.T) {
		ucMock := transaction.NewTransactionUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetTransactionByNumber", 332980303).Return([]model.TransactionDetail{}, fmt.Errorf("some error"))

		URL := "http://localhost:5000/transaction"
		request := httptest.NewRequest(http.MethodGet, URL, nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetTransactionByNumber(c)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("get transaction by number failed : error string conversion to integer", func(t *testing.T) {
		ucMock := transaction.NewTransactionUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		transactionNumber := "asadad"
		URL := fmt.Sprintf("http://localhost:5000/transaction?transaction_id=%s", transactionNumber)
		request := httptest.NewRequest(http.MethodGet, URL, nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetTransactionByNumber(c)
		response := recorder.Result()
	
		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("get transaction by number failed : id must be > 0", func(t *testing.T) {
		ucMock := transaction.NewTransactionUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetTransactionByNumber", 332980303).Return([]model.TransactionDetail{}, fmt.Errorf("id must be > 0"))
		
		transactionNumber := "332980303"
		URL := fmt.Sprintf("http://localhost:5000/transaction?transaction_id=%s", transactionNumber)
		request := httptest.NewRequest(http.MethodGet, URL, nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetTransactionByNumber(c)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("get transaction by number failed", func(t *testing.T) {
		ucMock := transaction.NewTransactionUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetTransactionByNumber", 332980303).Return([]model.TransactionDetail{}, fmt.Errorf("some error"))
		
		transactionNumber := "332980303"
		URL := fmt.Sprintf("http://localhost:5000/transaction?transaction_id=%s", transactionNumber)
		request := httptest.NewRequest(http.MethodGet, URL, nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetTransactionByNumber(c)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	// t.Run("create bulk transaction success", func(t *testing.T) {
	// 	ucMock := transaction.NewTransactionUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	voucherCode := "Ph1ncon"

	// 	req := model.TransactionDetailBulkRequest{
	// 		Items: []model.TransactionDetailItemRequest{
	// 			{
	// 				Item: "Tumbler_Phincon",
	// 				Quantity: 3,
	// 			},
	// 			{
	// 				Item:"Kaos_Phincon_2",
	// 				Quantity: 5,
	// 			},
	// 			{
	// 				Item: "Lanyard_Phincon_2",
	// 				Quantity: 3,
	// 			},
	// 		},
	// 		Name: "Utsman",
	// 		Pay: 1000000,
	// 	}

	// 	ucMock.On("CreateBulkTransactionDetail", voucherCode, req).Return(
	// 		[]model.TransactionDetail{
	// 			{
	// 				Id: 0,
	// 				Item: "Tumbler_Phincon",
	// 				Price: 30000,
	// 				Quantity: 3,
	// 				Total: 90000,
	// 				Transaction: model.Transaction{
	// 					Id: 0,
	// 					TransactionNumber: 332980303,
	// 					Name: "Utsman",
	// 					Quantity: 11,
	// 					Discount: 0.3,
	// 					Total: 336000,
	// 					Pay: 1000000,
	// 				},
	// 			},
	// 			{
	// 				Id: 0,
	// 				Item: "Kaos_Phincon_2",
	// 				Price: 30000,
	// 				Quantity: 5,
	// 				Total: 150000,
	// 				Transaction: model.Transaction{
	// 					Id: 0,
	// 					TransactionNumber: 332980303,
	// 					Name: "Utsman",
	// 					Quantity: 11,
	// 					Discount: 0.3,
	// 					Total: 336000,
	// 					Pay: 1000000,
	// 				},
	// 			},
	// 			{
	// 				Id: 0,
	// 				Item: "Lanyard_Phincon_2",
	// 				Price: 80000,
	// 				Quantity: 3,
	// 				Total: 240000,
	// 				Transaction: model.Transaction{
	// 					Id: 0,
	// 					TransactionNumber: 332980303,
	// 					Name: "Utsman",
	// 					Quantity: 11,
	// 					Discount: 0.3,
	// 					Total: 336000,
	// 					Pay: 1000000,
	// 				},
	// 			},
	// 		}, nil)


	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)

	// 	URL := fmt.Sprintf("http://localhost:5000/transaction?voucher_code=%s", voucherCode)
	// 	request := httptest.NewRequest(http.MethodPost, URL, body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.CreateBulkTransactionDetail(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusCreated, response.StatusCode)
	// })
	
	// t.Run("create bulk transaction failed : json decoder", func(t *testing.T) {
	// 	ucMock := transaction.NewTransactionUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	voucherCode := "Ph1ncon"
	// 	body := bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0})
		
	// 	URL := fmt.Sprintf("http://localhost:5000/transaction?voucher_code=%s", voucherCode)
	// 	request := httptest.NewRequest(http.MethodPost, URL, body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.CreateBulkTransactionDetail(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	// })

	// t.Run("create bulk transaction failed : quantity transaction should not be negative", func(t *testing.T) {
	// 	ucMock := transaction.NewTransactionUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	voucherCode := "Ph1ncon"
	// 	req := model.TransactionDetailBulkRequest{
	// 		Items: []model.TransactionDetailItemRequest{
	// 			{
	// 				Item: "Tumbler_Phincon",
	// 				Quantity: 3,
	// 			},
	// 			{
	// 				Item:"Kaos_Phincon_2",
	// 				Quantity: 5,
	// 			},
	// 			{
	// 				Item: "Lanyard_Phincon_2",
	// 				Quantity: 3,
	// 			},
	// 		},
	// 		Name: "Utsman",
	// 		Pay: 1000000,
	// 	}

	// 	ucMock.On("CreateBulkTransactionDetail", voucherCode, req).Return([]model.TransactionDetail{}, fmt.Errorf("quantity transaction should not be negative"))

	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)
		
	// 	URL := fmt.Sprintf("http://localhost:5000/transaction?voucher_code=%s", voucherCode)
	// 	request := httptest.NewRequest(http.MethodPost, URL, body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.CreateBulkTransactionDetail(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	// })

	// t.Run("create bulk transaction failed", func(t *testing.T) {
	// 	ucMock := transaction.NewTransactionUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	voucherCode := "Ph1ncon"
	// 	req := model.TransactionDetailBulkRequest{
	// 		Items: []model.TransactionDetailItemRequest{
	// 			{
	// 				Item: "Tumbler_Phincon",
	// 				Quantity: 3,
	// 			},
	// 			{
	// 				Item:"Kaos_Phincon_2",
	// 				Quantity: 5,
	// 			},
	// 			{
	// 				Item: "Lanyard_Phincon_2",
	// 				Quantity: 3,
	// 			},
	// 		},
	// 		Name: "Utsman",
	// 		Pay: 1000000,
	// 	}

	// 	ucMock.On("CreateBulkTransactionDetail", voucherCode, req).Return([]model.TransactionDetail{}, fmt.Errorf("some error"))
		
	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)
		
	// 	URL := fmt.Sprintf("http://localhost:5000/transaction?voucher_code=%s", voucherCode)
	// 	request := httptest.NewRequest(http.MethodPost, URL, body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.CreateBulkTransactionDetail(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	// })
}