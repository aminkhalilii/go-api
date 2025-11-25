package messagebroker

type MessageBrokerInterface interface {
	DeclareExchange(name, kind string) error
	DeclareQueue(name string) error
	BindQueue(queue, exchange, routingKey string) error
	Publish(exchange, routingKey string, message []byte) error
	Consume(queue string) (<-chan []byte, error)
}
