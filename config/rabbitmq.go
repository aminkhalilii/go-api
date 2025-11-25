package config

import (
	"go-api/internal/messageBroker/rabbitmq"
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQClient *rabbitmq.RabbitMQ // exported
var rabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func InitRabbitMQ() {
	var err error
	rabbitConn, err = amqp.Dial(AppConfig.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %v", err)
	}

	RabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}

	// init wrapper
	RabbitMQClient = rabbitmq.NewRabbitMQ(RabbitChannel)

	log.Println("RabbitMQ connected")

}

// CloseRabbitMQ closes the RabbitMQ connection
func CloseRabbitMQ() {
	if RabbitChannel != nil {
		RabbitChannel.Close()
	}
	if rabbitConn != nil {
		rabbitConn.Close()
	}
}
