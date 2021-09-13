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
		false,   // durable : A durable queue only means that the queue definition will survive a server restart, not the messages in it
		false,   // delete when unused : if set to false, the messages will be lost during the server restart 
		false,   // exclusive : Exclusive queues are only accessible by the connection that declares them and will be deleted when the connection closes
		false,   // no-wait : When noWait is true, the queue will assume to be declared on the server
		nil,     // arguments : No extra arguments provided 
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",      // exchange : When you want a single message to be delivered to a single queue, you can publish to the default exchange with the routingKey of the queue name. 
		q.Name,  // routing key : When you want a single message to be delivered to a single queue, you can publish to the default exchange with the routingKey of the queue name. 
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