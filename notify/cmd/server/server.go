package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/idirall22/crypto_app/notify/service/model"
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

	// q, err := ch.QueueDeclare(
	// 	"email", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	q2, err := ch.QueueDeclare(
		"notification", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// data, err := json.Marshal(model.RegisterUserConfirmationEmailParams{
	// 	Email:     "idirall22@gmail.com",
	// 	FirstName: "idir",
	// 	Subject:   "simple hello",
	// 	Body:      "thank you",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data2, err := json.Marshal(model.Notification{
		UserID:    1,
		Type:      "notif",
		Title:     "transaction done",
		Content:   "you just send money",
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}

	// body := []byte(data)
	// err = ch.Publish(
	// 	"",     // exchange
	// 	q.Name, // routing key
	// 	false,  // mandatory
	// 	false,  // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        []byte(body),
	// 	})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	body := []byte(data2)
	err = ch.Publish(
		"",      // exchange
		q2.Name, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
