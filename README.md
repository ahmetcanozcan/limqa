<h1 align="center">LiMQA</h1>

<div align="center">
 
</div>
<div align="center">
  <strong>abstraction for AMQP</strong>
</div>
<div align="center">
  a <code>Go</code> library that makes easier to <code> AMQP</code> communication
</div>

<br/>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [About](#about)
- [Installation](#installation)
- [Usage](#usage)
  - [Hello World](#hello-world)
  - [Options](#options)
- [CLI](#cli)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## About

Limqa is an client abstraction for AMQP communication using Consumer and Producer design. Limqa is built top of [streadway/amqp](https://github.com/streadway/amqp).

## Installation

You can install using go get

```bash
go get github.com/ahmetcanozcan/limqa
```

## Usage

### Hello World

```go
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
 // If you are sure that exchange is declared before
 // you don't have to declare it again.
 // consumer can be instantiated without DeclareExchange flag:
 // consumer, _ := limqa.NewConsumer(base,"_queue","_exchange")

 producer, _ := limqa.NewProducer(base,"_exchange")

 // Produce a message
 producer.Produce([]byte("Hello World"))

 // Get message from the consumer
 msg := consumer.Consume()

 fmt.Println(string(msg))
 // Output : Hello World
}

```

### Options

a consumer or producer can be configured using limqa options

```go
// ...
consumer1, _ := limqa.NewConsumer(base,"_queue","_exchange",limqa.DeclareExchange(true),limqa.NoAck(false))
// ...
consumer2, _ := limqa.NewConsumer(base,"_queue2","_exchange",limqa.NoWait(true),limqa.NoAck(true),limqa.NoLocal(true))

producer, _  := limqa.NewProducer(base,"_exchange",limqa.Durable(true),limqa.AutoDelete(false))
// ...
```

## CLI

Limqa can be used by cli for message producing or consuming using Limqa library. It can be installed using `go get`

```bash
go get github.com/ahmetcanozcan/limqa/limqa
```

Now, we can produce and consume messages.

```powershell
limqa produce -m "Hello World" -exchange hello_world
# Output : message sent
```

```powershell
limqa consume -queue helo_queue_1 -e hello_world
# Output : Hello World
```

for more help for flags and commands use:

```powershell
limqa -help
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Ahmetcan ÖZCAN - [email](mailto:ahmetcanozcan7@gmail.com)
