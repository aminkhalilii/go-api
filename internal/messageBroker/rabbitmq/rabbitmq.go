package rabbitmq

import "github.com/streadway/amqp"

// RabbitMQ implements message broker behavior using RabbitMQ.
type RabbitMQ struct {
	Channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ instance.
func NewRabbitMQ(ch *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{
		Channel: ch,
	}
}

// DeclareExchange creates an exchange if it does not exist.
func (r *RabbitMQ) DeclareExchange(name, kind string) error {
	return r.Channel.ExchangeDeclare(
		name,
		kind,  // direct, fanout, topic, headers
		true,  // durable
		false, // auto-deleted
		false,
		false,
		nil,
	)
}

// DeclareQueue creates a queue if it does not exist.
func (r *RabbitMQ) DeclareQueue(name string) error {
	_, err := r.Channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// BindQueue binds a queue to an exchange using a routing key.
func (r *RabbitMQ) BindQueue(queue, exchange, routingKey string) error {
	return r.Channel.QueueBind(
		queue,
		routingKey,
		exchange,
		false,
		nil,
	)
}

// Publish sends a message to the exchange with a routing key.
func (r *RabbitMQ) Publish(exchange, routingKey string, message []byte) error {
	return r.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}

// Consume receives messages from a queue.
func (r *RabbitMQ) Consume(queue string) (<-chan []byte, error) {
	msgs, err := r.Channel.Consume(
		queue,
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

	out := make(chan []byte)

	go func() {
		for msg := range msgs {
			out <- msg.Body
		}
	}()

	return out, nil
}
