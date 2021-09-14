package main

import (
	"os"
	"log"
	"strings"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {
	body := bodyForm(os.Args)
	log.Println(body)

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
		"hello", // queue name 
		true,    // to make the queue persist even after rabbitmq restart (need to do it on both consumer and provider level)
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Unable to declare a queue")

	err = ch.Publish(
		"",  // exchange 
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // to make the undelivered messages persist even after rabbit mq restarts 
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)

}

// helper function
func bodyForm(args []string) string {
	var s string 
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

// helper function
func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}