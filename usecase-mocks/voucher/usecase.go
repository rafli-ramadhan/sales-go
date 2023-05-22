package voucher

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct{
	mock.Mock
}

func NewVoucherUsecaseMock() *UsecaseMock {
	return &UsecaseMock{}
}

func (m *UsecaseMock) GetList() (response []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	response = ret.Get(0).([]model.Voucher)
	err = ret.Error(1)
	return response, err
}

func (m *UsecaseMock) GetVoucherByCode(name string) (response model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	response = ret.Get(0).(model.Voucher)
	err = ret.Error(1)
	return response, err
}

func (m *UsecaseMock) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	response = ret.Get(0).([]model.Voucher)
	err = ret.Error(1)
	return response, err
}