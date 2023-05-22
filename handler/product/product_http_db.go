package product

import (
	"encoding/json"
	"net/http"
	"sales-go/helpers/rest"
	"sales-go/model"
	"sales-go/usecase/product"
)

type dbhttphandler struct {
	usecase product.ProductUseCase
}

func NewDBHTTPHandler(usecase product.ProductUseCase) *dbhttphandler {
	return &dbhttphandler{
		usecase: usecase,
	}
}

func (handler *dbhttphandler) GetList(w http.ResponseWriter, r *http.Request) {
	res, err := handler.usecase.GetList()
	if err != nil {
		rest.ResponseError(w, r, http.StatusInternalServerError, err)
		return
	}
	rest.ResponseData(w, r, http.StatusOK, res)
}

func (handler *dbhttphandler) Create(w http.ResponseWriter, r *http.Request) {
	req := []model.ProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rest.ResponseError(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := handler.usecase.Create(req)
	if err != nil {
		if err.Error() == "name should not be empty" || err.Error() == "price should be > 0" {
			rest.ResponseError(w, r, http.StatusBadRequest, err)
			return
		} else {
			rest.ResponseError(w, r, http.StatusInternalServerError, err)
			return
		}
	}
	rest.ResponseData(w, r, http.StatusOK, res)
}
