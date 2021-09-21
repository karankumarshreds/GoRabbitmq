package main

import "log"

func (app *App) ErrorListener() {
	msgs, err := app.listener.Listen("sub_4_payment_queue", "direct_logs", "error")
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
}