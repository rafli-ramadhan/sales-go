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
	listVoucher = m.Called().Get(0).([]model.Voucher)
	return listVoucher, nil
}

func (m *RepoMock) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	voucherData = m.Called(code).Get(0).(model.Voucher)
	return voucherData, nil
}

func (m *RepoMock) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	response = m.Called(req).Get(0).([]model.Voucher)
	return response, nil
}
