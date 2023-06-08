package product

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

type ClientPostgresql struct {
	suite.Suite
	db *sql.DB
	mock sqlmock.Sqlmock
	repo Repositorier
}

func TestRepoPostgresql(t *testing.T) {
	suite.Run(t, new(ClientPostgresql))
}

func (c *ClientPostgresql) SetupTest() {
	var err error
	c.db, c.mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	c.repo = NewPostgreSQLHTTPRepository(c.db)
}

func (c *ClientPostgresql) SetupSuite() {
	log.Println("set up test     : before all test executed")
}

func (c *ClientPostgresql) TearDownTest() {
	log.Println("tear down test  : after each test executed")
}

func (c *ClientPostgresql) TearDownSetup() {
	log.Println("tear down setup : after all test executed")
}

func (c *ClientPostgresql) AfterTest() {
	log.Println("AFTER TEST 	 : after all test executed")
}

func (c *ClientPostgresql) TestGetListContactSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Kaos Phincon", 30000).
		AddRow(2, "Lanyard_Phincon", 30000)

	c.mock.ExpectPrepare(`SELECT id, name, price FROM product`).
		WillBeClosed().ExpectQuery().WillReturnRows(rows)

	res, err := c.repo.GetList()
	if err != nil {
		c.T().Error(err)
	}

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientPostgresql) TestGetListContactFailPrepareStmt() {
	c.mock.ExpectPrepare(`SELECT id, name, price FROM product`).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestGetListContactFailQuery() {
	c.mock.ExpectPrepare(`SELECT id, name, price FROM product`).
		WillBeClosed().
		ExpectQuery().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetList()

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestGetProductByNameSuccess() {
	row := sqlmock.NewRows([]string{"id", "code", "persen"}).
		AddRow(1, "Kaos_Phincon", 30000)

	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, name, price FROM product WHERE name = $1`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Kaos_Phincon").
		WillReturnRows(row)

	res, err := c.repo.GetProductByName("Kaos_Phincon")
	if err != nil {
		c.T().Error(err)
	}

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientPostgresql) TestGetProductByNameFailPrepareStmt() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, name, price FROM product WHERE name = $1`)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetProductByName("Kaos_Phincon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestGetProductByNameFailQuery() {
	c.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, name, price FROM product WHERE name = $1`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Kaos_Phincon").
		WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.GetProductByName("Kaos_Phincon")

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestCreateSuccess() {
	var price1 float64 = 30000
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Kaos_Phincon", price1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).AddRow("1", "Kaos_Phincon", 30000))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.ProductRequest{
		{
			Name:  "Kaos_Phincon",
			Price: 30000,
		},
	})

	require.NoError(c.T(), err)
	require.NotEmpty(c.T(), res)
}

func (c *ClientPostgresql) TestCreateFailedBeginTransaction() {
	c.mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))

	res, err := c.repo.Create([]model.ProductRequest{
		{
			Name:  "Kaos_Phincon",
			Price: 30000,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestCreateFailedPrepareStmt() {
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`)).
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.ProductRequest{
		{
			Name:  "Kaos_Phincon",
			Price: 30000,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}

func (c *ClientPostgresql) TestCreateFailedQuery() {
	var price1 float64 = 30000
	c.mock.ExpectBegin()
	c.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`)).
		WillBeClosed().
		ExpectQuery().
		WithArgs("Kaos_Phincon", price1).
		WillReturnError(fmt.Errorf("some error"))
	c.mock.ExpectCommit()

	res, err := c.repo.Create([]model.ProductRequest{
		{
			Name:  "Kaos_Phincon",
			Price: 30000,
		},
	})

	require.Error(c.T(), err)
	require.Empty(c.T(), res)
}
