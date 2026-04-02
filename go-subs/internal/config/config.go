package config

import "os"

type Config struct {
	MQTTBroker string
	RabbitURL  string
}

func Load() Config {
	return Config{
		MQTTBroker: getEnv("MQTT_BROKER", "tcp://mqtt:1883"),
		RabbitURL:  getEnv("RABBIT_URL", "amqp://guest:guest@rabbitmq:5672/"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
