package dbInterface

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func connect() *gorm.DB {
	db, err := gorm.Open("postgres",
		"dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		println(err.Error())
		panic("Failed to connect to db! ")
	}
	return db
}

var Psql = connect()
