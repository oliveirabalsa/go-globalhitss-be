package config

import (
	"log"
	"os"

	"github.com/oliveirabalsa/go-globalhitss-be/db"
	"github.com/oliveirabalsa/go-globalhitss-be/queue"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupConsumer(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	queueName := os.Getenv("RABBITMQ_QUEUE")
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil

}

func InitServices() (*amqp.Channel, *amqp.Connection, *gorm.DB) {
	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	conn, ch, err := queue.NewRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	return ch, conn, db
}
