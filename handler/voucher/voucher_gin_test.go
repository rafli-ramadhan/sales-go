package voucher

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
	"sales-go/usecase-mocks/voucher"
)

func TestHandlerGin(t *testing.T) {
	t.Run("test handler get list voucher success", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Voucher{
			{
				Id: 1,
				Code: "Ph1ncon",
				Persen: 30,
			},
			{
				Id: 2,
				Code: "Phintraco",
				Persen: 20,
			},
		}, nil)

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/voucher", nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetList(c)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})
	
	t.Run("test handler get list voucher failed", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Voucher{}, fmt.Errorf("some error"))

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/voucher", nil)
		recorder := httptest.NewRecorder()		
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetList(c)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

// 	t.Run("test handler create multiple voucher success", func(t *testing.T) {
// 		ucMock := voucher.NewVoucherUsecaseMock()
// 		handler := NewGinDBHTTPHandler(ucMock)

// 		req := []model.VoucherRequest{
// 			{
// 				Code: "Ph1ncon",
// 				Persen: 30,
// 			},
// 			{
// 				Code: "Phintraco",
// 				Persen: 20,
// 			},
// 		}
// 		ucMock.On("Create", req).Return([]model.Voucher{
// 			{
// 				Id: 1,
// 				Code: "Ph1ncon",
// 				Persen: 30,
// 			},
// 			{
// 				Id: 2,
// 				Code: "Phintraco",
// 				Persen: 20,
// 			},
// 		}, nil)

// 		jsonByte, err := json.Marshal(req)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		body := bytes.NewReader(jsonByte)

// 		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
// 		recorder := httptest.NewRecorder()

// 		handler.Create(recorder, request)
// 		response := recorder.Result()

// 		require.Equal(t, http.StatusOK, response.StatusCode)
// 	})

// 	t.Run("test handler create multiple voucher failed : json decoder", func(t *testing.T) {
// 		ucMock := voucher.NewVoucherUsecaseMock()
// 		handler := NewGinDBHTTPHandler(ucMock)

// 		body := bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0})

// 		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
// 		recorder := httptest.NewRecorder()

// 		handler.Create(recorder, request)
// 		response := recorder.Result()

// 		require.Equal(t, http.StatusBadRequest, response.StatusCode)
// 	})

// 	t.Run("test handler create multiple voucher failed with some empty code", func(t *testing.T) {
// 		ucMock := voucher.NewVoucherUsecaseMock()
// 		handler := NewGinDBHTTPHandler(ucMock)

// 		req := []model.VoucherRequest{
// 			{
// 				Code: "",
// 				Persen: 30,
// 			},
// 			{
// 				Code: "Phintraco",
// 				Persen: 20,
// 			},
// 		}
// 		ucMock.On("Create", req).Return([]model.Voucher{}, fmt.Errorf("code should not be empty"))

// 		jsonByte, err := json.Marshal(req)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		body := bytes.NewReader(jsonByte)

// 		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
// 		recorder := httptest.NewRecorder()

// 		handler.Create(recorder, request)
// 		response := recorder.Result()

// 		require.Equal(t, http.StatusBadRequest, response.StatusCode)
// 	})

// 	t.Run("test handler create multiple voucher failed", func(t *testing.T) {
// 		ucMock := voucher.NewVoucherUsecaseMock()
// 		handler := NewGinDBHTTPHandler(ucMock)

// 		req := []model.VoucherRequest{
// 			{
// 				Code: "Ph1ncon",
// 				Persen: 30,
// 			},
// 			{
// 				Code: "Phintraco",
// 				Persen: 20,
// 			},
// 		}
// 		ucMock.On("Create", req).Return([]model.Voucher{}, fmt.Errorf("some error"))

// 		jsonByte, err := json.Marshal(req)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		body := bytes.NewReader(jsonByte)

// 		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
// 		recorder := httptest.NewRecorder()

// 		handler.Create(recorder, request)
// 		response := recorder.Result()

// 		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
// 	})
}