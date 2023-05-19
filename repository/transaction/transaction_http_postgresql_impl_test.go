package transaction

import (
	"database/sql"
	"fmt"
	"log"
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

func (c *Client) TestGetTransactionByNumber() {
	row := sqlmock.AddRow().
	c.mock.ExpectPrepare(`SELECT id, transaction_number, name, quantity, discount, total, pay FROM transaction WHERE transaction_number = $1`).
		WillBeClosed().
		ExpectQuery().
		WillReturnRows(row)

}
