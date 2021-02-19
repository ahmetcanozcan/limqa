package limqa_test

import (
	"fmt"
	"testing"

	"github.com/ahmetcanozcan/limqa"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)


func TestNewProducer(t *testing.T) {

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
			err: &amqp.Error{
				Code:406,
				Reason:"PRECONDITION_FAILED - inequivalent arg 'internal' for exchange 'test_producer' in vhost '/': received 'true' but current is 'false'",
				Server:true,
				Recover:true},
			opts: []limqa.Option{limqa.Durable(true),limqa.Internal(true)},
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
		_,err  := limqa.NewProducer(ti.base,"test_producer",ti.opts...)
		assert.Equal(t,err,ti.err,fmt.Sprintf("case #%d",i))
	}
}


func TestProducer_Produce(t *testing.T) {
	p,err := limqa.NewProducer(pbase(),"test_producer_2")
	
	assert.Nil(t,err)
	assert.NotNil(t,p)

	err = p.Produce([]byte(""))
	assert.Nil(t,err)

}