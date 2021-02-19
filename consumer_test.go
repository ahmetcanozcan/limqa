package limqa_test

import (
	"fmt"
	"testing"

	"github.com/ahmetcanozcan/limqa"
	"github.com/stretchr/testify/assert"
)


func TestNewConsumer(t *testing.T) {
	table := []struct{
		base *limqa.Base
		err error
		opts []limqa.Option
	}{
		{
			base: pbase(),
			err: nil,
			opts: []limqa.Option{},
		},
		{
			base: pbase(),
			err: nil,
			opts: []limqa.Option{limqa.AutoAck(false),limqa.Durable(true),limqa.NoLocal(true)},
		},
		{
			base: pbase(),
			err: nil,
			opts: []limqa.Option{limqa.Durable(true),limqa.NoWait(true)},
		},
		{
			base: nil,
			err: limqa.ErrNilBase,
			opts: []limqa.Option{},
		},
		{
			base: limqa.New(),
			err : limqa.ErrBaseNotConnected,
			opts: []limqa.Option{},
		},
	}
	for i, ti := range table {
		_,err  := limqa.NewConsumer(ti.base,fmt.Sprintf("%s-%d",_queue,i),_exchange,ti.opts...)
		assert.Equal(t,err,ti.err,fmt.Sprintf("case #%d",i))
	}
}



func TestConsumer_Consume(t *testing.T) {
	const msg = "TEST MESSAGE"
	b := pbase()
	p,err := limqa.NewProducer(b,_exchange)
	assert.Nil(t,err)

	err = p.Produce([]byte(msg));
	assert.Nil(t,err)
	
	c, err := limqa.NewConsumer(b,_queue,_exchange)
	assert.Nil(t,err)

	m := c.Consume()
	assert.Equal(t,msg,string(m))

}