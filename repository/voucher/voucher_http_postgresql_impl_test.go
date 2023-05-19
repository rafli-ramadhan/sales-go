package voucher

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Client struct {
	suite.Suite
	db *sql.DB
	mock sqlmock.Sqlmock
	repo Repositorier
}

func TestRepoVoucher(t *testing.T) {
	suite.Run(t, new(Client))
}

func (c *Client) SetupTest() {
	var err error
	c.db, c.mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	c.repo = NewPostgreSQLHTTPRepository(c.db)
}

func (c *Client) SetupSuite() {
	log.Println("set up test     : before all test executed")
}

func (c *Client) TearDownTest() {
	log.Println("tear down test  : after each test executed")
}

func (c *Client) TearDownSetup() {
	log.Println("tear down setup : after all test executed")
}

func (c *Client) AfterTest() {
	log.Println("after test 	 : after all test executed")
}

func (c *Client) TestGetListContactSuccess() {
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

func (c *Client) TestGetListContactFailPrepareStmt() {
	c.mock.ExpectPrepare(`SELECT id, name, price FROM voucher`).
		WillBeClosed().WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *Client) TestGetListContactFailQuery() {
	c.mock.ExpectPrepare(`SELECT id, name, price FROM voucher`).
		WillBeClosed().ExpectQuery().WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *Client) TestGetVoucherByCodeSuccess() {
	row := sqlmock.NewRows([]string{"id", "code", "persen"}).
		AddRow("1", "Ph1ncon", "30")

	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = $1`)).
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

func (c *Client) TestGetVoucherByCodeFailStmt() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = $1`)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetVoucherByCode("Ph1ncon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *Client) TestGetVoucherByCodeFailQuery() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, code, persen FROM voucher WHERE code = $1`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Ph1ncon").
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetVoucherByCode("Ph1ncon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}