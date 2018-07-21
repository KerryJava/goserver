package base

import (
	"database/sql"
	"github.com/KerryJava/goserver/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *sql.DB = NewMysql()
var OrmDB *gorm.DB = NewOrmDB()

func NewOrmDB() *gorm.DB {

	log.Println("DSN:", config.Content.DSN)
	db, err := gorm.Open("mysql", config.Content.DSN)

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return db
}

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

func init() {
	log.Println("init db")
}
