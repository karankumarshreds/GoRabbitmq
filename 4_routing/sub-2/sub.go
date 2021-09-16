/*

		LISTENS TO ROUTING KEY "info" PASSED IN ARGUMENTS BY PUBLISHER

*/

package main

import (
	"log"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {
	
	// dial a connection
	conn, err := amqp.Dial(connString)
	failOnError(err, "Unable to dial a connection")

	// open a channel
	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")

	// declare a queue 
	q, err := ch.QueueDeclare(
		"",          // name of the queue 
		false,       // durable 
		false,       // autoDelete 
		true,        // exclusive 
		false,       // noWait
		nil,
	)
	failOnError(err, "Unable to declare a queue")
	
	// declare an exchange 
	err = ch.ExchangeDeclare(
		"direct_logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Unable to declare exchange")

	// bind the exchange 
	err = ch.QueueBind(
		q.Name,           // name of the queue 
		"info",           // name of the key
		"direct_logs",    // name of the exchange 
		false,
		nil,
	)
	failOnError(err, "Unable to bind the exchange")


	// listen for the messages 
	msgs, err := ch.Consume(
		q.Name,
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Unable to consume messages")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

// helper functions 

func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}