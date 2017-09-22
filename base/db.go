package base

import (
	"database/sql"
	"log"
)
import "github.com/KerryJava/goserver/config"
import _ "github.com/go-sql-driver/mysql"

var DB *sql.DB = NewMysql()

func NewMysql() *sql.DB {
	log.Println("DSN:", config.Content.DSN)
	db, err := sql.Open("mysql", config.Content.DSN)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	db.SetMaxIdleConns(config.Content.MaxIdleConns)
	db.SetMaxOpenConns(config.Content.MaxOpenConns)
	return db
}
