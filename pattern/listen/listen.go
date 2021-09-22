package listen 

import (
	"log"
	"bytes"
	"encoding/json"
	"github.com/streadway/amqp"
)

type Listener struct {
	Channel *amqp.Channel
	Msgs <-chan amqp.Delivery
}

// New will create a new instance of listener 
func New(ch *amqp.Channel) Listener{
	return Listener{
		Channel: ch,
	}
}

func (l *Listener) Listen(queue string, exchange string, topic string) (err error) {
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
	
	msgs, err := l.Channel.Consume(
		q.Name,
		"",    // consumer name (keeping it unique for identity of listener)
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait 
		nil,   // args 
	)
	
	l.Msgs = msgs
	
	return err
}


type ErrorMessage struct {
	Name string `json:"Name"`
	Age string  `json:"Age"`
}

type Callback func(interface{}) error
type MessageStruct interface{}
type CBInterface interface {
	ErrorCallback(errorMessage ErrorMessage)
}

func (l *Listener) OnMessage(callback CBInterface, msgStruct MessageStruct) error {
	switch msgStruct.(type) {
	case ErrorMessage:
		var errorMessage ErrorMessage
		go func() {	
			for d := range l.Msgs {
				decoder := json.NewDecoder(bytes.NewReader(d.Body))
				err := decoder.Decode(&errorMessage)
				if err != nil {
					log.Println("Error while decoding", err)
				}
				callback.ErrorCallback(errorMessage)
			}
		}()
	default:
		return nil
	}
	return nil
}

// func (l *Listener) iterateMessages(messageStructInstance interface{},ch *chan ErrorMessage) {
// 	go func() {	
// 		for d := range l.Msgs {
// 			decoder := json.NewDecoder(bytes.NewReader(d.Body))
// 			err := decoder.Decode(&messageStructInstance)
// 			if err != nil {
// 				log.Println("Error while decoding", err)
// 			}
// 			ch <- messageStructInstance
// 		}
// 	}()
// }

func returnIfError(err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}  else {
		return nil, nil
	}
}




