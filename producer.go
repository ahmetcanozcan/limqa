package limqa

import (
	"github.com/streadway/amqp"
)


var _ Producer = (*producer)(nil)

var (
	defaultProducerOpt = []Option{ Durable(true), AutoDelete(false),Internal(false),NoWait(false) }
)

// Producer :
type Producer interface {
	Produce(data []byte) error
}

type producer struct {
	base *Base
	exchange string
}

// NewProducer :
func NewProducer(base *Base, name string,options ...Option) (Producer,error) {
	if base == nil {
		return nil,ErrNilBase
	}

	if !base.IsConnected() {
		return nil,ErrBaseNotConnected
	}
	
	opts := parseOptions(defaultProducerOpt,options)
	f := getFlags(opts)

	if err := base.ch.ExchangeDeclare(name,"topic",f[durable] , f[autoDelete], f[internal], f[noWait], nil); err != nil {
		return nil,err
	}

	return &producer{
		base: base,
		exchange: name,
	},nil
}

func (p *producer) Produce(b []byte) error {
	m := amqp.Publishing {
		Body: b,
	}
	return produceMessage(p.base.ch,p.exchange,m)
}



func produceMessage(ch *amqp.Channel,exchange string,msg amqp.Publishing) error {
	return ch.Publish(exchange,"random-key",false,false,msg)
}