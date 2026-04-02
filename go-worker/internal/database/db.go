package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) *gorm.DB {

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}
