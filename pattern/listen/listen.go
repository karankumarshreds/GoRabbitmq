package listen 

import (
	"github.com/streadway/amqp"
)

type Listener struct {
	Channel *amqp.Channel
}

// New will create a new instance of listener 
func New(ch *amqp.Channel) Listener{
	return Listener{
		Channel: ch,
	}
}

func (l Listener) Listen(queue string, exchange string, topic string) (msgs <-chan amqp.Delivery, err error) {
	q, err := l.Channel.QueueDeclare(
		queue,   // queue name 
		true,    // durable 
		false,   // autoDelete
		false,   // exclusive
		false,   // noWait 
		nil,
	)
	returnIfError(err)

	err = l.Channel.ExchangeDeclare(
		exchange,  // exchange name 
		"direct",  // type
		true,      // durable 
		false,     // autoDelete
		false,     // internal 
		false,     // noWait 
		nil,
	)
	returnIfError(err)

	err = l.Channel.QueueBind(
		q.Name,
		topic,
		exchange,
		false,
		nil,
	)
	returnIfError(err)
	
	msgs, err = l.Channel.Consume(
		q.Name,
		"",    // consumer name (keeping it unique for identity of listener)
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait 
		nil,   // args 
	)
	returnIfError(err)

	return msgs, nil	
	
}

func returnIfError(err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}  else {
		return nil, nil
	}
}




