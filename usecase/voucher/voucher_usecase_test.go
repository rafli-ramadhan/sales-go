package voucher

import (
	"testing"
	"github.com/stretchr/testify/require"

	"sales-go/mocks/voucher"
	"sales-go/model"
)

func TestGetList(t *testing.T) {
	t.Run("get list voucher", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		mockSuccess.On("GetList").Return([]model.Voucher{
			{
				Id: 1,
				Code: "ph1ncon",
				Persen: 20,
			},
			{
				Id: 2,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 3,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 4,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 5,
				Code: "Ph1nc0n",
				Persen: 30,
			},
			{
				Id: 6,
				Code: "ph1ncon2",
				Persen: 20,
			},
		})

		_, err := usecase.GetList()
		require.NoError(t, err)
	})
}

func TestGetVoucherByCode(t *testing.T) {
	t.Run("test get voucher by code", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		code := "Ph1ncon"
		mockSuccess.On("GetVoucherByCode", code).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		})

		_, err := usecase.GetVoucherByCode(code)
		require.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("test create voucher", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		req := []model.VoucherRequest{
			{
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		mockSuccess.On("Create", req).Return([]model.Voucher{
			{
				Id: 4,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 5,
				Code: "Ph1nc0n",
				Persen: 30,
			},
		})

		_, err := usecase.Create(req)
		require.NoError(t, err)
	})
}
