package voucher

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"sales-go/model"
	"time"
)

type repositoryhttpgormpostgresql struct {
	db *gorm.DB
}

func NewPostgreSQLGormHTTPRepository(db *gorm.DB) *repositoryhttpgormpostgresql {
	return &repositoryhttpgormpostgresql{
		db: db,
	}
}

func (repo *repositoryhttpgormpostgresql) GetList() (listVoucher []model.Voucher, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT id, code, persen FROM voucher`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}

	for res.Next() {
		var voucher model.Voucher
		res.Scan(&voucher.Id, &voucher.Code, &voucher.Persen)
		listVoucher = append(listVoucher, voucher)
	}

	return
}

func (repo *repositoryhttpgormpostgresql) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT id, code, persen FROM voucher WHERE code = $1`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, code)
	if err != nil {
		return
	}

	for res.Next() {
		res.Scan(&voucherData.Id, &voucherData.Code, &voucherData.Persen)
	}
	return
}

func (repo *repositoryhttpgormpostgresql) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `INSERT INTO voucher (code, persen) VALUES ($1, $2) RETURNING id, code, persen`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	for _, v := range req {
		var temp model.Voucher
		err = stmt.QueryRowContext(ctx, v.Code, v.Persen).Scan(&temp.Id, &temp.Code, &temp.Persen)
		if err != nil {
			trx.Rollback()
			return []model.Voucher{}, err
		}
		response = append(response, temp)
	}

	trx.Commit()

	return
}
