package limqa

import "github.com/streadway/amqp"

var (
	defaultConsumerOpt = []Option{DeclareExchange(false),AutoAck(true),Exclusive(false),NoLocal(false)}
)

// Consumer :
type Consumer interface {
	Consume() []byte
}

type consumer struct {
	q    amqp.Queue
	base *Base
	consumeCh <-chan amqp.Delivery
}

// NewConsumer :
func NewConsumer(base *Base, queue, exchange string, options ...Option) (Consumer, error) {
	if base == nil {
		return nil, ErrNilBase
	}

	if !base.IsConnected() {
		return nil, ErrBaseNotConnected
	}

	def := append(defaultProducerOpt, defaultConsumerOpt...)
	opts := parseOptions(def, options)
	f := getFlags(opts)

	if f[declareExchange] {
		// Make sure exchange is declared
		err := base.ch.ExchangeDeclare(exchange, "topic", f[durable], f[autoDelete], f[internal], f[noWait], nil)
		if err != nil {
			return nil, err
		}
	}

	// Bind queue and exchange with wildcard '#'
	// For more information : https://www.rabbitmq.com/tutorials/tutorial-five-go.html
	q, err := base.ch.QueueDeclare(queue, f[durable], f[autoDelete], f[exclusive], f[noWait], nil)
	if err != nil {
		return nil, err
	}

	err = base.ch.QueueBind(queue, "#",exchange,f[noWait],nil)

	if err != nil {
		return nil,err
	}

	cch ,err := base.ch.Consume(queue,"",f[autoAck],f[exclusive],f[noLocal],f[noWait],nil)

	return &consumer {
		q: q,
		base: base,
		consumeCh: cch,
	}, nil

}


func (c *consumer) Consume() []byte {
	return (<- c.consumeCh).Body
}