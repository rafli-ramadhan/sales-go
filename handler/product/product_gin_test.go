package product

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
	"sales-go/usecase-mocks/product"
)

func TestHandlerGin(t *testing.T) {
	t.Run("test handler get list product success", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Product{
			{
				Id: 1,
				Name: "Kaos_Phincon_2",
				Price: 30000,
			},
			{
				Id: 2,
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Id: 3,
				Name: "Lanyard_Phincon_2",
				Price: 80000,
			},
		}, nil)

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/product", nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetList(c)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})
	
	t.Run("test handler get list product failed", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewGinDBHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Product{}, fmt.Errorf("some error"))

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/product", nil)
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request = request

		handler.GetList(c)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	// t.Run("test handler create multiple product success", func(t *testing.T) {
	// 	ucMock := product.NewProductUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	req := []model.ProductRequest{
	// 		{
	// 				Name: "Kaos_Phincon",
	// 				Price: 30000,
	// 		},
	// 		{
	// 				Name: "Lanyard_Phincon",
	// 				Price: 80000,
	// 		},
	// 	}
	// 	ucMock.On("Create", req).Return([]model.Product{
	// 		{
	// 			Id: 1,
	// 			Name: "Kaos_Phincon",
	// 			Price: 30000,
	// 		},
	// 		{
	// 			Id: 2,
	// 			Name: "Lanyard_Phincon",
	// 			Price: 80000,
	// 		},
	// 	}, nil)

	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)
	// 	fmt.Println("BODY : ", body)

	// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
	// 	recorder := httptest.NewRecorder()
	// 	c, _ := gin.CreateTestContext(recorder)
	// 	c.Request = request
	// 	fmt.Println("c Request : ", c.Request)

	// 	handler.Create(c)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusOK, response.StatusCode)
	// })

	// t.Run("test handler create multiple product failed : json decoder", func(t *testing.T) {
	// 	ucMock := product.NewProductUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	body := bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0})

	// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.Create(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	// })

	// t.Run("test handler create multiple product failed with some empty name", func(t *testing.T) {
	// 	ucMock := product.NewProductUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	req := []model.ProductRequest{
	// 		{
	// 				Name: "",
	// 				Price: 30000,
	// 		},
	// 		{
	// 				Name: "Lanyard_Phincon",
	// 				Price: 80000,
	// 		},
	// 	}
	// 	ucMock.On("Create", req).Return([]model.Product{}, fmt.Errorf("name should not be empty"))

	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)

	// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.Create(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	// })

	// t.Run("test handler create multiple product failed", func(t *testing.T) {
	// 	ucMock := product.NewProductUsecaseMock()
	// 	handler := NewGinDBHTTPHandler(ucMock)

	// 	req := []model.ProductRequest{
	// 		{
	// 				Name: "Kaos_Phincon",
	// 				Price: 30000,
	// 		},
	// 		{
	// 				Name: "Lanyard_Phincon",
	// 				Price: 80000,
	// 		},
	// 	}
	// 	ucMock.On("Create", req).Return([]model.Product{}, fmt.Errorf("some error"))

	// 	jsonByte, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(jsonByte)

	// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.Create(recorder, request)
	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	// })
}