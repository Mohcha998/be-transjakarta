package publisher

import (
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Start() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt:1883")
	client := mqtt.NewClient(opts)

	token := client.Connect()
	token.Wait()

	for {
		data := map[string]interface{}{
			"vehicle_id": "B1234XYZ",
			"latitude":   -6.2088,
			"longitude":  106.8456,
			"timestamp":  time.Now().Unix(),
		}

		payload, _ := json.Marshal(data)

		log.Println("Publising:", string(payload))

		client.Publish("/fleet/vehicle/B1234XYZ/location", 0, false, payload)

		time.Sleep(2 * time.Second)
	}
}
