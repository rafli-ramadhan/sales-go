package transaction

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
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo Repositorier
}

func TestRepoTransaction(t *testing.T) {
	suite.Run(t, new(Client))
}

func (c *Client) SetupTest() {
	var err error
	c.db, c.mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err.Error()))
	}

	c.repo = NewPostgreSQLHTTPRepository(c.db)
}

func (c *Client) SetupSuite() {
	log.Println("setup suite     : executed before all of test")
}

func (c *Client) TearDownTest() {
	log.Println("setup test      : executed before each of test")
}

func (c *Client) TearDownSetup() {
	log.Println("tear down setup : executed after all of test")
}

func (c *Client) AfterTest() {
	log.Println("after test      : executed after all of test")
}

func (c *Client) TestGetTransactionByNumberSuccess() {
	query1 := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`
	
	row1 := sqlmock.NewRows([]string{"id", "transaction_number", "name", "quantity", "discount", "total", "pay"}).
		AddRow(1, 288029617, "Utsman", 11, 0, 480000, 1000000)
	row2 := sqlmock.NewRows([]string{"id", "item", "price", "quantity", "total"}).
		AddRow(1, "Tumber_Phincon", 30000, 3, 90000).
		AddRow(2, "Kaos_Phincon", 30000, 3, 90000).
		AddRow(3, "Lanyard_Phincon", 30000, 3, 90000)

	c.mock.ExpectPrepare(regexp.QuoteMeta(query1)).
		WillBeClosed().
		ExpectQuery().
		WillReturnRows(row1)

	c.mock.ExpectPrepare(regexp.QuoteMeta(query2)).
		WillBeClosed().
		ExpectQuery().
		WillReturnRows(row2)

	res, err := c.repo.GetTransactionByNumber(288029617)
	if err != nil {
		c.T().Error(err)
	}

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *Client) TestGetTransactionByNumberFailPrepareStmt() {
	query1 := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`

	c.mock.ExpectPrepare(regexp.QuoteMeta(query1)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	c.mock.ExpectPrepare(regexp.QuoteMeta(query2)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetTransactionByNumber(288029617)

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *Client) TestGetTransactionByNumberFailQuery() {
	query1 := `SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`
	query2 := `SELECT id, item, price, quantity, total FROM transaction_detail WHERE transaction_id = $1`

	c.mock.ExpectPrepare(regexp.QuoteMeta(query1)).
		WillBeClosed().
		ExpectQuery().
		WillReturnError(fmt.Errorf("some error"))

	c.mock.ExpectPrepare(regexp.QuoteMeta(query2)).
		WillBeClosed().
		ExpectQuery().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetTransactionByNumber(288029617)

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}