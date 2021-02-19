package main

import (
	"fmt"

	"github.com/ahmetcanozcan/limqa"
)



func main() {
	uri := "amqp://guest:guest@localhost:5672"
	base := limqa.New()
	
	base.Connect(uri)
	
	consumer, _ := limqa.NewConsumer(base,"_queue","_exchange",limqa.DeclareExchange(true))
	
	producer, _ := limqa.NewProducer(base,"_exchange")
	
	// Produce message
	producer.Produce([]byte("Hello World"))

	// Get message by consumer
	msg := consumer.Consume()

	fmt.Println(string(msg))
	// Output : Hello World
}