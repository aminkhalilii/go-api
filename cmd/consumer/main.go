package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:secret123@localhost:5672/")
	if err != nil {
		log.Fatalf("error connection %v", err)
	}

	defer conn.Close()

	fmt.Println("connected rabbit successfuly")

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("error channel %v", err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"user_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌  error  Exchange: %v", err)
	}

	q, err := ch.QueueDeclare("myQueue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ error   Queue: %v", err)
	}

	err = ch.QueueBind(q.Name, "user.registered", "user_exchange", false, nil)
	if err != nil {
		log.Fatalf("❌ error  Bind  : %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("send error   : %v", err)
	}

	fmt.Println("listening to messages")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("📩 message has dalivered: %s\n", d.Body)
		}
	}()
	<-forever
}
