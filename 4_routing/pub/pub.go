package main

import (
	"os"
	"log"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {

	// dial a connection
	conn, err := amqp.Dial(connString)
	failOnError(err, "Unable to establish a connection")

	// open a channel 
	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")

	// declare an exchange with its type 
	err  = ch.ExchangeDeclare(
		"direct_logs",    // name of the exchange 
		"direct",         // type of exchange 
		true,             // durable 
		false,            // auto-delete 
		false,            // internal
		false,            // no-wait
		nil,
	)
	failOnError(err, "Unable to declare an exchange")

	log.Printf("Publishing the message to routing key === %v\n", routingKeyFromArgs(os.Args))
	// publish the message on the exchange 
	ch.Publish(
		"direct_logs",                        // name of exchange 
		routingKeyFromArgs(os.Args),          // routing key (info|warning|error)
		false,                                // mandatory
		false,                                // immediate 
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World!"),
		},
	)

}

// helper functions 

func routingKeyFromArgs(args []string) string {
	var s string
		if (len(args) < 2) || os.Args[1] == "" {
			s = "info"
		} else {
			s = os.Args[1]
		}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}