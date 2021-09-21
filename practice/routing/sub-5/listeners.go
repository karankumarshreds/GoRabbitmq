package main

import (
	// "fmt"
	"log"
	"bytes"
	"encoding/json"
)


type ErrorMessage struct {
	Name string `json:"Name"`
	Age string  `json:"Age"`
}

func (app *App) ErrorListener() {
	var errorMessage ErrorMessage 
	msgs, err := app.listener.Listen("sub-5-queue", "direct_logs", "error")
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			decoder := json.NewDecoder(bytes.NewReader(d.Body))
			err = decoder.Decode(&errorMessage)
			if err != nil {
				log.Println("Error while decoding", err)
			}
			log.Printf("Got Person %v With Age %v", errorMessage.Name, errorMessage.Age)
		}
	}()


	// make a function for listener package OnMessage
	// it will take an argument of a function (which will take arg of type struct of message) and return an error 
	// and the logic that we will run
	// we will return the same error/nil if we get any
  em := ErrorMessage{}
	onMessage(localLogic, &em)
}



type Callback func() error
type MessageStruct struct{}

func onMessage(callback Callback, messageStructPointer interface{}) error {
	
	return nil
}
// actual callback function
func localLogic() error {
	log.Println("Do something in the db")
	return nil
}


