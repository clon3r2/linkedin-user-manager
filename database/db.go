package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var er error

func InitializeDatabase() {
	DBConn, er = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if er != nil {
		log.Fatal(er)
	}
}
