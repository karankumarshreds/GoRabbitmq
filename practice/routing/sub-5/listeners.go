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
	app.listener.OnMessage(app, listen.ErrorMessage{})
}

func (app *App) ErrorCallback(em listen.ErrorMessage) {
	app.SaveErrorToDb(em)
}


func (app *App) SaveErrorToDb(em listen.ErrorMessage) {
		log.Printf("Got Person %v With Age %v", em.Name, em.Age)
}
