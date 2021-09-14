package main 

import (
	"log"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {


	// create a connection
	conn, err := amqp.Dial(connString) 
	failOnError(err, "Unable to establish connection")
	defer conn.Close()

	// open a channel 
	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")
	defer ch.Close()

	// declare a queue 
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Unable to declare a queue")

	// publish on the queue 
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World"),
		},
	)
	failOnError(err, "Failed to publish message")


}

func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}