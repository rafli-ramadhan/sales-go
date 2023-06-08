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
	// get return value dari mock
	response = ret.Get(0).([]model.Voucher)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *UsecaseMock) GetVoucherByCode(name string) (response model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	// get return value dari mock
	response = ret.Get(0).(model.Voucher)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *UsecaseMock) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	// get return value dari mock
	response = ret.Get(0).([]model.Voucher)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}