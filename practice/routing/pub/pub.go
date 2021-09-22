package main

import (
	// "fmt"
	"os"
	"log"
	"encoding/json"
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
	type ErrorMessage struct {
		Name string 
		Age string
	}
	msg := ErrorMessage{
		Name: "Karan Kumar",
		Age: "27",
	}

	body, _ := json.Marshal(msg)
	
	ch.Publish(
		"direct_logs",                        // name of exchange 
		routingKeyFromArgs(os.Args),          // routing key (info|warning|error)
		false,                                // mandatory
		false,                                // immediate 
		amqp.Publishing{
			ContentType: "text/plain",
			Body: body,
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