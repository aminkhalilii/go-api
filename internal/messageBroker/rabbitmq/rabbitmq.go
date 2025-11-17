package rabbitmq

import "github.com/streadway/amqp"

// RabbitMQ implements message broker behavior using RabbitMQ.
type RabbitMQ struct {
	url     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ instance.
func NewRabbitMQ(url string) *RabbitMQ {
	return &RabbitMQ{
		url: url,
	}
}

// Connect establishes the connection and channel.
func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	r.channel = ch
	return nil
}

// DeclareExchange creates an exchange if it does not exist.
func (r *RabbitMQ) DeclareExchange(name, kind string) error {
	return r.channel.ExchangeDeclare(
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
	_, err := r.channel.QueueDeclare(
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
	return r.channel.QueueBind(
		queue,
		routingKey,
		exchange,
		false,
		nil,
	)
}

// Publish sends a message to the exchange with a routing key.
func (r *RabbitMQ) Publish(exchange, routingKey string, message []byte) error {
	return r.channel.Publish(
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
	msgs, err := r.channel.Consume(
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

// Close closes the channel and connection.
func (r *RabbitMQ) Close() error {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	return nil
}
