package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GormMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
