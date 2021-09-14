package main 

import (
	"time"
	"bytes"
	"log"
	"github.com/streadway/amqp"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {

	// create connection 
	conn, _ := amqp.Dial(connString)

	// open a channel
	ch, _ := conn.Channel()

	// declare a queue from which you want to consume 
	q, _ := ch.QueueDeclare(
		"hello",
		true,    // durable : to make the queue persist even after rabbitmq restart 
		         // (need to do it on both consumer and provider level)
		false,
		false,
		false,
		nil,
	)
	
	// consume from the channel 
	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// create a go channel
	forever := make(chan bool)

	// iterate over the returned go subroutine 
	go func() {
		for d := range msgs {
			log.Printf("Received message %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	<-forever
	
}


// func failOnError(err error, msg string) {
// 	if err != nil {
//     log.Fatalf("%s: %s", msg, err)
//   }
// }