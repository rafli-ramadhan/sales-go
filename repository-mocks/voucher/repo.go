package voucher

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct{
	mock.Mock
}

func NewVoucherRepoMock() *RepoMock {
	return &RepoMock{}
}


func (m *RepoMock) GetList() (listVoucher []model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	// get return value dari mock
	listVoucher = ret.Get(0).([]model.Voucher)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *RepoMock) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(code)
	// get return value dari mock
	voucherData = ret.Get(0).(model.Voucher)
	if ret.Get(1) != nil {
		// type assertion
		err = ret.Get(1).(error)
	}
	return
}

func (m *RepoMock) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
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
