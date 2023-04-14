package client

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"sales-go/config"
	logger "sales-go/helpers/logging"
)

var (
	_ = godotenv.Load()
	conDB 				= config.NewConfig()
	connString string 	= ""
	Database 			= os.Getenv("DATABASE")
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
	if dbOpt.Database == "mysql" {
		// format : "username:password@tcp(host:port)/database_name"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v", conDB.MySQL.Username, conDB.MySQL.Password, conDB.MySQL.Host, conDB.MySQL.Port, conDB.MySQL.Database)
	}

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
		
	logger.Infof(fmt.Sprintf("Running mysql on %s on port %s\n", conDB.MySQL.Host, conDB.MySQL.Port))
	
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}