package database

import (
	"log"
	"time"

	"go-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) *gorm.DB {

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("DB Connected")
			break
		}

		log.Println("Waiting DB...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = db.AutoMigrate(&models.VehicleLocation{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}
