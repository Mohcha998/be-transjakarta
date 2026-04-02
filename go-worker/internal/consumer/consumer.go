package consumer

import (
	"encoding/json"
	"log"
	"time"

	"go-worker/internal/models"

	"gorm.io/gorm"
)

func Start(db *gorm.DB, url string) {

	ch := connectRabbitWithRetry(url)

	err := ch.ExchangeDeclare(
		"fleet.events",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"geofence_alerts",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		q.Name,
		"",
		"fleet.events",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Worker consuming geofence events...")

	for m := range msgs {

		var d map[string]interface{}
		err := json.Unmarshal(m.Body, &d)
		if err != nil {
			log.Println("invalid json")
			m.Nack(false, false)
			continue
		}

		vehicleID, ok := d["vehicle_id"].(string)
		if !ok {
			log.Println("invalid vehicle_id")
			m.Nack(false, false)
			continue
		}

		event, ok := d["event"].(string)
		if !ok {
			log.Println("invalid event")
			m.Nack(false, false)
			continue
		}

		location, ok := d["location"].(map[string]interface{})
		if !ok {
			log.Println("invalid location")
			m.Nack(false, false)
			continue
		}

		lat, ok := location["latitude"].(float64)
		if !ok {
			log.Println("invalid latitude")
			m.Nack(false, false)
			continue
		}

		lon, ok := location["longitude"].(float64)
		if !ok {
			log.Println("invalid longitude")
			m.Nack(false, false)
			continue
		}

		tsFloat, ok := d["timestamp"].(float64)
		if !ok {
			log.Println("invalid timestamp")
			m.Nack(false, false)
			continue
		}

		ts := time.Unix(int64(tsFloat), 0)

		err = db.Create(&models.VehicleLocation{
			VehicleID: vehicleID,
			Latitude:  lat,
			Longitude: lon,
			Timestamp: ts,
		}).Error

		if err != nil {
			log.Println("DB error:", err)
			m.Nack(false, false)
			continue
		}

		log.Println("GEOFENCE EVENT")
		log.Println("Vehicle:", vehicleID)
		log.Println("Event:", event)
		log.Println("Location:", lat, lon)
		log.Println("Time:", ts)

		m.Ack(false)
	}
}
