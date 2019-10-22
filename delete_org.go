package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	args := amqp.Table{}

	args["x-dead-letter-exchange"] = "cramstack_dlx"

	_, err = ch.QueueDeclare(
		"warehouse_q", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		args,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	bytes, _ := json.Marshal(struct {
		Id   uint   `json:"id"`
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	}{1, "abcd", "cramstack"})

	err = ch.Publish(
		"warehouse_x",         // exchange
		"organization.delete", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         bytes,
		})
	failOnError(err, "Failed to publish a message")
}
