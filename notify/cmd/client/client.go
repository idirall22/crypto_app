package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://user:password@0.0.0.0:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	wait := make(chan bool)

	go func() {
		d, err := ch.Consume("hello", "", true, false, false, false, amqp.Table{})
		if err != nil {
			log.Fatal(err)
		}
		for {
			fmt.Println("waiting")
			result := <-d
			fmt.Println(string(result.Body))
		}
	}()
	<-wait
}
