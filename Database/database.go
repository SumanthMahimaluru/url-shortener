package database

import (
	"database/sql"
	"fmt"
	"time"
	"urlshortener/common"

	_ "github.com/go-sql-driver/mysql"
)

var client sql.DB

func Connect() error {
	user := common.Db.User
	host := common.Db.Host
	pwd := common.Db.Password
	port := common.Db.Port
	dbname := common.Db.DbName
	source := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	fmt.Println("database connected successfully")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	client = *db
	return nil
}

func GetClient() *sql.DB { return &client }
