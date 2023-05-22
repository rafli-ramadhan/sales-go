package voucher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"sales-go/model"
	"sales-go/usecase-mocks/voucher"
)

func TestHandlerHTTPJson(t *testing.T) {
	t.Run("test handler get list voucher success", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

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

		handler.GetList(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})
	
	t.Run("test handler get list voucher failed", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		ucMock.On("GetList").Return([]model.Voucher{}, fmt.Errorf("some error"))

		request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/voucher", nil)
		recorder := httptest.NewRecorder()

		handler.GetList(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple voucher success", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.VoucherRequest{
			{
				Code: "Ph1ncon",
				Persen: 30,
			},
			{
				Code: "Phintraco",
				Persen: 20,
			},
		}

		ucMock.On("GetVoucherByCode", "Ph1ncon").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
			
		ucMock.On("GetVoucherByCode", "Phintraco").Return(model.Voucher{}, fmt.Errorf("voucher not found"))

		ucMock.On("Create", req).Return([]model.Voucher{
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

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)
	})

	t.Run("test handler create multiple voucher failed : json decoder", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		body := bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0})

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple voucher failed with some empty code", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.VoucherRequest{
			{
				Code: "",
				Persen: 30,
			},
			{
				Code: "Phintraco",
				Persen: 20,
			},
		}

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test handler create multiple voucher failed with some negative persen ", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.VoucherRequest{
			{
				Code: "Phincon",
				Persen: -1,
			},
			{
				Code: "Phintraco",
				Persen: 20,
			},
		}

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test handler create multiple voucher failed : create error", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.VoucherRequest{
			{
				Code: "Ph1ncon",
				Persen: 30,
			},
			{
				Code: "Phintraco",
				Persen: 20,
			},
		}

		ucMock.On("GetVoucherByCode", "Ph1ncon").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
			
		ucMock.On("GetVoucherByCode", "Phintraco").Return(model.Voucher{}, fmt.Errorf("voucher not found"))

		ucMock.On("Create", req).Return([]model.Voucher{}, fmt.Errorf("some error"))

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test handler create multiple voucher failed : voucher exist", func(t *testing.T) {
		ucMock := voucher.NewVoucherUsecaseMock()
		handler := NewJsonHTTPHandler(ucMock)

		req := []model.VoucherRequest{
			{
				Code: "Ph1ncon",
				Persen: 30,
			},
			{
				Code: "Phintraco",
				Persen: 20,
			},
		}

		ucMock.On("GetVoucherByCode", "Ph1ncon").Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 30,
		} , nil)
			
		ucMock.On("GetVoucherByCode", "Phintraco").Return(model.Voucher{
			Id: 2,
			Code: "Phintraco",
			Persen: 20,
		}, nil)

		jsonByte, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonByte)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/voucher", body)
		recorder := httptest.NewRecorder()

		handler.Create(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusConflict, response.StatusCode)
	})
}