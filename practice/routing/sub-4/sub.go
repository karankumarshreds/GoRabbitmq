package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/karankumarshreds/GoRabbitmq/pattern/listen"
)

const connString = "amqp://guest:guest@localhost:5672/"

type App struct{
	listener *listen.Listener
}

func main() {

	var app App

	conn, err := amqp.Dial(connString)
	failOnError(err, "Unable to dial a connection")

	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel")

	l := listen.New(ch)
	app.listener = &l

	// Add listeners here
	// ===========================
	forever := make(chan bool)
	app.ErrorListener()
	<-forever
	// ===========================

}



func failOnError(err error, msg string) {
	if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}