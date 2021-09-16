package main

import (
	"log"
	"github.com/streadway/amqp"
)

func main() {
	// create connection 
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Unable to establish connection")

	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")

	q, err := ch.QueueDeclare(
		"",      // we dont need to specify name for our consumer queue as the publisher 
		         // exchange type is fanout which will broadcast the message to all the queues 
		false,   // durable 
		false,   // delete when unused 
		true,    // exclusive (only accessible by connections declaring them and are deleted once this connection closes)
		false,   // no-wait 
		nil,     // arguments 
	)
	failOnError(err, "Unable to declare queue")

	// now declare the exchange with the same name that was declared by the publisher 
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// bind the exchange with your queue to get messages from it 
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key left blank because the exchange is of type fanout 
		"logs", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

// helper function 
func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}