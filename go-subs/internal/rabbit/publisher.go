package rabbit

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Publisher struct {
	ch *amqp.Channel
}

func New(url string) *Publisher {

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

	return &Publisher{ch: ch}
}

func (p *Publisher) Publish(data interface{}) {

	body, err := json.Marshal(data)
	if err != nil {
		log.Println("json error:", err)
		return
	}

	err = p.ch.Publish(
		"fleet.events",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("publish error:", err)
		return
	}

	log.Println("📤 Sent to RabbitMQ")
}

func connectRabbitWithRetry(url string) *amqp.Channel {

	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {

		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("RabbitMQ Connected")
			break
		}

		log.Println("Retry connect RabbitMQ...")
		time.Sleep(3 * time.Second)
	}

	if conn == nil {
		log.Fatal("Failed connect RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return ch
}
