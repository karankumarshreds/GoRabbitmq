package main

import (
	// "fmt"
	"log"
	_"bytes"
	_"encoding/json"
	"github.com/karankumarshreds/GoRabbitmq/pattern/listen"
)


func (app *App) ErrorListener() {
	// var errorMessage ErrorMessage 
	err := app.listener.Listen("sub-5-queue", "direct_logs", "error")
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	errCh := make(chan listen.ErrorMessage)
	// app.listener.OnMessage()
	app.listener.OnMessage(&errCh, listen.ErrorMessage{})
	go processErrorMessages(&errCh)
	for d := range errCh {
		app.SaveErrorToDb(d)
	}
}



func (app *App) SaveErrorToDb(em ErrorMessage) {
		log.Printf("Got Person %v With Age %v", d.Name, d.Age)
}


// func processErrorMessages(ch *chan ErrorMessage) error {
	
// 	return nil
// }


// 👉 different types of messages(struct) 

