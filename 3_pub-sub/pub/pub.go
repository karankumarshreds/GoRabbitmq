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

	// let us create a "fanout" type exchange and call it "logs"
	ch.ExchangeDeclare(
		"logs",    // name of exchange 👈
		"fanout",  // type of exchange 
		true,      // durable (exchange will persist if rabbitmq restarts)
		false,     // auto-deleted (if true, it will delete the exchange if there are no bindings with the queues)
		false,     // internal (google)
		false,     // no-wait (When noWait is true, declare without waiting for a confirmation from the server)
		nil,       // arguments 
	)

	body := "Hello World!"

	err = ch.Publish(
		"logs",     // sending the message to exchange 👈 (which will further publish to queue based on exchange type)
		"",         // routing key (kept blank in case of fanout exchanges as we want to broadcast the messsage)
		false,      // mandatory (publishings can be undeliverable if set to true)
		false,      // immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)
	failOnError(err, "Error while publishing")
	
}

// helper function 
func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}