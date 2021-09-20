package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/karankumarshreds/GoRabbitmq/pattern/listen"
)

const connString = "amqp://guest:guest@localhost:5672/"

func main() {
	
	conn, err := amqp.Dial(connString)
	failOnError(err, "Unable to dial a connection")

	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")

	listener := listen.New(ch, "sub_4")
	
	msgs, err := listener.Listen(
		"sub_4_queue",
		"direct_logs",
		"error",
	)
	failOnError(err, "Unable to listen")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}