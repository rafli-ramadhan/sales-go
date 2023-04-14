package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sales-go/config"
	"time"
)

const (
	mysql = "mysql"
)

var (
	conDB = config.NewConfig()
	connString string = ""
)

type dbOption struct {
	Database string
}

func NewConnection(database string) dbOption {
	return dbOption{
		Database: database,
	}
}

func (dbOpt dbOption) GetMysqlConnection() (db *sql.DB) {
	if dbOpt.Database == mysql {
		// "username:password@tcp(host:port)/database_name"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v", conDB.MySQL.Username, conDB.MySQL.Password, conDB.MySQL.Host, conDB.MySQL.Port, conDB.MySQL.Database)
	}

	db, err := sql.Open(mysql, connString)
	if err != nil {
		panic(err)
	}
		
	log.Printf("Running mysql on %s on port %s\n", conDB.MySQL.Host, conDB.MySQL.Port)

	
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}