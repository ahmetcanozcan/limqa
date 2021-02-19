package limqa_test

import (
	"errors"
	"os"
	"testing"

	"github.com/ahmetcanozcan/limqa"
	"github.com/stretchr/testify/assert"
)

const (
	_uri = "amqp://guest:guest@localhost:5672"
	_queue = "test_queue_c"
	_exchange = "test_exchange_c"
)

var (
	_producer limqa.Producer
)


func TestMain(m *testing.M) {
	var err error
	pbase() // Run for checking that function does not panic
	_producer,err = limqa.NewProducer(pbase(),_exchange)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}





/* Test Utility Function */

func pbase() *limqa.Base{
	pbase := limqa.New()
	if err := pbase.Connect(_uri); err != nil {
		panic(err)
	}
	return pbase
}




func TestBase(t *testing.T) {

	table := []struct{
		uri string
		err error
	} {
		{
			uri: _uri,
			err: nil,
		},
		{
			uri:"invalid_uri",
			err : errors.New("AMQP scheme must be either 'amqp://' or 'amqps://'"),
		},
	}

	for _,ti := range table {
		b := limqa.New()
		err:=b.Connect(ti.uri)
	
		assert.Equal(t,err,ti.err)
		assert.Equal(t,b.IsConnected(),ti.err == nil)
	}



}