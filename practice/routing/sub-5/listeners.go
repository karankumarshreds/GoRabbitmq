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


	em := ErrorMessage{}
	onMessage(localLogic, em)

}

// make a function for listener package OnMessage
// it will take an argument of a function (which will take arg of type struct of message) and return an error 
// and the logic that we will run
// we will return the same error/nil if we get any
// em := ErrorMessage{}
// onMessage(localLogic, &em)


type Callback func() error
type MessageStruct interface{}

// func onMessage(callback Callback, msgStruct MessageStruct) error {

// 	go func() {
// 		for d := range msgs {
// 			log.Printf(" [x] %s", d.Body)
// 			decoder := json.NewDecoder(bytes.NewReader(d.Body))
// 			err = decoder.Decode(&errorMessage)
// 			if err != nil {
// 				log.Println("Error while decoding", err)
// 			}
// 			log.Printf("Got Person %v With Age %v", errorMessage.Name, errorMessage.Age)
// 		}
// 	}()

// 	// callback(&structPointer)
// 	switch msgStruct.(type) {
// 		case ErrorMessage: 
// 			log.Printf("Got Person %v With Age %v", msgStruct.Name, msgStruct.Age)
// 		return nil
// 	default: 
// 		return nil
// 	}
	
// }
// // actual callback function
func localLogic(em ErrorMessage) error {
	log.Printf("Got Person %v With Age %v", em.Name, em.Age)
	return nil
}
