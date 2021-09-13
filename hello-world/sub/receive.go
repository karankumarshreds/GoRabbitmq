package main 

import (
	"log"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/* declare the queue from which we are going to consume */
	/* since this queue has already been declared it won't be declared again */
	/* we also do this because we accidently might start consumer before the publisher */
	/* and we want to make sure that the queue exists before we try to consume messages from it */
	q, err := ch.QueueDeclare(
		"hello", // name of the queue from which we will be consuming
		false,   // durable : A durable queue only means that the queue definition will survive a server restart, not the messages in it
		false,   // delete when unused : if set to false, the messages will be lost during the server restart 
		false,   // exclusive : Exclusive queues are only accessible by the connection that declares them and will be deleted when the connection closes
		false,   // no-wait : When noWait is true, the queue will assume to be declared on the server
		nil,     // arguments : No extra arguments provided 
	)
	failOnError(err, "Failed to declare a queue")

	/* msgs returned here is a go routine */
	msgs, err := ch.Consume(
		q.Name,         // queue name 
		"",             // consumer (blank string will assign consumer a unique tag automatically)
		true,           // auto-ack
		false,          // exclusive (exclusive is true, the server will ensure that this is the sole consumer from this queue)
		false,          // no-local (not supported by RabbitMQ)
		false,          // no-wait (noWait is true, do not wait for the server to confirm the request and immediately begin deliveries)
		nil,             // no other arguments 
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received message %s", d.Body)
		}
	}()

	<-forever
}

// helper function
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}