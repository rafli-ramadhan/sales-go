package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"sales-go/model"
	"sales-go/usecase-mocks/product"
)

func TestHandlerHTTPJson(t *testing.T) {
	t.Run("test handler get list product success", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

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

		handler.GetList(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})
	
	t.Run("test handler get list product failed", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Product{}, fmt.Errorf("some error"))

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/product", nil)
		recorder := httptest.NewRecorder()

		handler.GetList(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple product success", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}

		ucMock.On("GetProductByName", "Kaos_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
			
		ucMock.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))

		ucMock.On("Create", req).Return([]model.Product{
			{
				Id: 1,
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Id: 2,
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}, nil)

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)
	})

	t.Run("test handler create multiple product : product exist", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}

		ucMock.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos_Phincon",
			Price: 30000,
		}, nil)
			
		ucMock.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)

		ucMock.On("Create", req).Return([]model.Product{
			{
				Id: 1,
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Id: 2,
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}, nil)

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusConflict, response.StatusCode)
	})

	t.Run("test handler create multiple product failed : json decoder", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		body := bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0})

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple product failed : some empty name", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
					Name: "",
					Price: 30000,
			},
			{
					Name: "Lanyard_Phincon",
					Price: 80000,
			},
		}

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test handler create multiple product failed : some negative price", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
					Name: "Kaos_Phincon",
					Price: -30000,
			},
			{
					Name: "Lanyard_Phincon",
					Price: 80000,
			},
		}

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test handler create multiple product failed : create error", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
					Name: "Kaos_Phincon",
					Price: 30000,
			},
			{
					Name: "Lanyard_Phincon",
					Price: 80000,
			},
		}

		ucMock.On("GetProductByName", "Kaos_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
			
		ucMock.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))

		ucMock.On("Create", req).Return([]model.Product{}, fmt.Errorf("some error"))

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple product success", func(t *testing.T) {
		ucMock := product.NewProductUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}

		ucMock.On("GetProductByName", "Kaos_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
			
		ucMock.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))

		ucMock.On("Create", req).Return([]model.Product{
			{
				Id: 1,
				Name: "",
				Price: 30000,
			},
			{
				Id: 2,
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
		}, nil)

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/product", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)
	})
}