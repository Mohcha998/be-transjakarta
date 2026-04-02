package main

import (
	"log"
	"net/http"

	"go-subs/internal/config"
	"go-subs/internal/mqtt"
	"go-subs/internal/rabbit"
)

func main() {

	cfg := config.Load()

	pub := rabbit.New(cfg.RabbitURL)

	go startHealth()

	mqtt.Start(cfg.MQTTBroker, pub)
}

func startHealth() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("Subs healthcheck running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
