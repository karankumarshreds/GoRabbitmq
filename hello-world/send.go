package main

import (
	"log"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {

	// The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication
	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	// Channel represents an AMQP channel. Used as a context for valid message exchange.
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")
	defer ch.Close()

	// To send, we must declare a queue for us to send to; then we can publish a message to the queue
	q, err := ch.QueueDeclare(
		"hello", // name of the queue 
		false,   // durable 
		false,   // delete when unused 
		false,   // exclusive 
		false,   // no-wait 
		nil,     // arguments 
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",      // exchange 
		q.Name,  // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World!"),
		},
	)
	failOnError(err, "Failed to publish a message")

}

// helper function
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}