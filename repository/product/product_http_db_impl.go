package product

import (
	"fmt"
	"context"
	"sales-go/db"
	"sales-go/model"
	"time"
)

type repositoryhttpdb struct {}

func NewDBHTTPRepository () *repositoryhttpdb {
	return &repositoryhttpdb{}
}

func (repo *repositoryhttpdb) GetList() (listProduct []model.Product, err error) {
	db := client.NewConnection("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}

	for res.Next() {
		var temp model.Product
		res.Scan(&temp.Id, &temp.Name, &temp.Price)
		fmt.Println(temp)

		listProduct = append(listProduct, temp)
	}

	return
}

func (repo *repositoryhttpdb) GetProductByName(name string) (productData model.Product, err error) {
	db := client.NewConnection("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, price FROM product WHERE name = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, name)
	if err != nil {
		return
	}

	for res.Next() {
		res.Scan(&productData.Id, &productData.Name, &productData.Price)
	}
	return
}

func (repo *repositoryhttpdb) Create(req []model.ProductRequest) (result []model.Product, err error) {
	db := client.NewConnection("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `INSERT INTO product (name, price) VALUES (?, ?)`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	
	for _, v := range req {
		res, err := stmt.ExecContext(ctx, v.Name, v.Price)
		if err != nil {
			trx.Rollback()
			return []model.Product{}, err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return []model.Product{}, err
		}

		result = append(result, model.Product{
			Id:    int(lastID),
			Name:  v.Name,
			Price: v.Price,
		})
	}

	trx.Commit()

	return
}