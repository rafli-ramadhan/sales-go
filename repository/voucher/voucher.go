package voucher

import (
	"errors"
	"sales-go/model"
)

type repository struct {}

func NewRepository() *repository {
	return &repository{}
}

type Repositorier interface {
	GetList() []model.Voucher
	GetVoucherByCode(code string) (voucherData model.Voucher, err error)
	Create(req model.VoucherRequest) model.Voucher
}

func (repo *repository) GetList() []model.Voucher {
	return model.VoucherSlice
}

func (repo *repository) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	for _, v := range model.VoucherSlice {
		if v.Code == code {
			voucherData = v
		}
	}

	emptyStruct := model.Voucher{}
	if voucherData == emptyStruct {
		err = errors.New("Voucher not found")
		return
	}
	return
}

func (repo *repository) Create(req model.VoucherRequest) model.Voucher {
	newData := model.Voucher{
		Id:     int64(len(model.VoucherSlice)) + 1,
		Code:   req.Code,
		Persen: req.Persen,
	}
	model.VoucherSlice = append(model.VoucherSlice, newData)

	return newData
}
