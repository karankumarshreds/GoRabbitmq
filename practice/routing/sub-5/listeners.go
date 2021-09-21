package main

import (
	"fmt"
	"log"
	"bytes"
	"encoding/json"
)

type ErrorMessage struct {
	Name string 
	Age string
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
			body, _ := json.Marshal(d.Body)
			fmt.Println("JSON BODY", body)
			decoder := json.NewDecoder(bytes.NewReader(body))
			err = decoder.Decode(&errorMessage)
			if err != nil {
				log.Println("Error while decoding", err)
			}
			log.Printf("Got Person %v With Age %v", errorMessage.Name, errorMessage.Age)
		}
	}()
	
}



