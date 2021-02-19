// Package limqa implements a
package limqa

import "github.com/streadway/amqp"

// Base :
type Base struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

// New :
func New() *Base {
	return &Base{}
}

// IsConnected :
func (b *Base) IsConnected() bool {
	return (b.conn != nil ) && (b.ch != nil ) && !b.conn.IsClosed()
}

// Connect :
func (b *Base) Connect(uri string) error {
	c,err := amqp.Dial(uri)
	if err != nil {
		return err
	}

	ch, err := c.Channel()
	if err != nil {
		return err
	}

	b.conn = c
	b.ch = ch
	return nil
}


