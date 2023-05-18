package config

import (
	"log"
	"time"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	mysql_drv "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:suryakant1234@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True"
	d, err := gorm.Open(mysql_drv.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		log.Fatal("Error", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
