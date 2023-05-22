package voucher

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"sales-go/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ClientMysql struct {
	suite.Suite
	db *sql.DB
	mock sqlmock.Sqlmock
	repo Repositorier
}

func TestRepoMysql(t *testing.T) {
	suite.Run(t, new(ClientMysql))
}

func (c *ClientMysql) SetupTest() {
	var err error
	c.db, c.mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	c.repo = NewMySQLHTTPRepository(c.db)
}

func (c *ClientMysql) SetupSuite() {
	log.Println("set up test     : before all test executed")
}

func (c *ClientMysql) TearDownTest() {
	log.Println("tear down test  : after each test executed")
}

func (c *ClientMysql) TearDownSetup() {
	log.Println("tear down setup : after all test executed")
}

func (c *ClientMysql) AfterTest() {
	log.Println("after test 	 : after all test executed")
}

func (c *ClientMysql) TestGetListContactSuccess() {
	rows := sqlmock.NewRows([]string{"id", "code", "persen"}).
		AddRow("1", "Ph1ncon", "30").
		AddRow("2", "Phintraco", "20")

	c.mock.ExpectPrepare(`SELECT id, code, persen FROM voucher`).
		WillBeClosed().ExpectQuery().WillReturnRows(rows)

	res, err := c.repo.GetList()
	if err != nil {
		c.T().Error(err)
	}

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientMysql) TestGetListContactFailPrepareStmt() {
	c.mock.ExpectPrepare(`SELECT id, code, persen FROM voucher`).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestGetListContactFailQuery() {
	c.mock.ExpectPrepare(`SELECT id, code, persen FROM voucher`).
		WillBeClosed().
		ExpectQuery().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestGetVoucherByCodeSuccess() {
	row := sqlmock.NewRows([]string{"id", "code", "persen"}).
		AddRow("1", "Ph1ncon", "30")

	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = ?`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Ph1ncon").
		WillReturnRows(row)

	res, err := c.repo.GetVoucherByCode("Ph1ncon")
	if err != nil {
		c.T().Error(err)
	}

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientMysql) TestGetVoucherByCodeFailPrepareStmt() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = ?`)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetVoucherByCode("Ph1ncon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestGetVoucherByCodeFailQuery() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = ?`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Ph1ncon").
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetVoucherByCode("Ph1ncon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestCreateSuccess() {
	var persen1 float64 = 30
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO voucher (code, persen) values (?, ?)`)).
		WillBeClosed().
		ExpectExec().
		WithArgs("Ph1ncon", persen1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.VoucherRequest{
		{
			Code:  "Ph1ncon",
			Persen: 30,
		},
	})

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientMysql) TestCreateFailedBeginTransaction() {
	c.mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.Create([]model.VoucherRequest{
		{
			Code:  "Ph1ncon",
			Persen: 30,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestCreateFailedPrepareStmt() {
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO voucher (code, persen) values (?, ?)`)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.VoucherRequest{
		{
			Code:  "Ph1ncon",
			Persen: 30,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestCreateFailedQuery() {
	var persen1 float64 = 30
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO voucher (code, persen) values (?, ?)`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Ph1ncon", persen1).
		WillReturnError(fmt.Errorf("some error"))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.VoucherRequest{
		{
			Code:  "Ph1ncon",
			Persen: 30,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientMysql) TestCreateFailedLastIDIsMinus() {
	var persen1 float64 = 30
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO voucher (code, persen) values (?, ?)`)).
		WillBeClosed().
		ExpectExec().
		WithArgs("Ph1ncon", persen1).
		WillReturnResult(sqlmock.NewResult(-1, 1))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.VoucherRequest{
		{
			Code:  "Ph1ncon",
			Persen: 30,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}
