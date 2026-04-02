package main

import (
	"log"
	"net/http"

	"go-worker/internal/config"
	"go-worker/internal/consumer"
	"go-worker/internal/database"

	"gorm.io/gorm"
)

func main() {

	cfg := config.Load()

	db := database.Init(cfg.DB)

	go startHealthServer(db)

	consumer.Start(db, cfg.Rabbit)
}

func startHealthServer(dbConn interface{}) {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		sqlDB, err := dbConn.(*gorm.DB).DB()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("db error"))
			return
		}

		if err := sqlDB.Ping(); err != nil {
			w.WriteHeader(500)
			w.Write([]byte("db down"))
			return
		}

		w.Write([]byte("ok"))
	})

	log.Println("Healthcheck running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
