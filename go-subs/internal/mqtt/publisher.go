package mqtt

import (
	"encoding/json"
	"log"
	"time"

	"go-subs/internal/geofence"
	"go-subs/internal/rabbit"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Payload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func Start(broker string, pub *rabbit.Publisher) {

	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(5 * time.Second)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT error:", token.Error())
	}

	log.Println("MQTT Connected")

	if token := client.Subscribe("/fleet/vehicle/+/location", 0, func(c mqtt.Client, m mqtt.Message) {

		var p Payload
		err := json.Unmarshal(m.Payload(), &p)
		if err != nil {
			log.Println("invalid json")
			return
		}

		if p.VehicleID == "" {
			log.Println("invalid vehicle_id")
			return
		}

		log.Println("MQTT:", p)

		if geofence.Check(p.Latitude, p.Longitude) {

			// log.Println("GEOFENCE TRIGGERED:", p.VehicleID)

			pub.Publish(map[string]interface{}{
				"vehicle_id": p.VehicleID,
				"event":      "geofence_entry",
				"location": map[string]interface{}{
					"latitude":  p.Latitude,
					"longitude": p.Longitude,
				},
				"timestamp": p.Timestamp,
			})
		}

	}); token.Wait() && token.Error() != nil {
		log.Fatal("Subscribe error:", token.Error())
	}

	select {}
}
