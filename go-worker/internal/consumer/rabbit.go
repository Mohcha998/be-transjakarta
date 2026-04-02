package consumer

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func connectRabbitWithRetry(url string) *amqp.Channel {

	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {

		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("RabbitMQ Connected (Worker)")
			break
		}

		log.Println("Retry connect RabbitMQ (Worker)...")
		time.Sleep(3 * time.Second)
	}

	if conn == nil {
		log.Fatal("Worker gagal connect RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return ch
}
