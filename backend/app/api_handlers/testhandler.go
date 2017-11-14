package api_handlers

import (
	"net/http"
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"log"
)

func TestHandler(w http.ResponseWriter, _ *http.Request) {
	SendData();
	fmt.Fprintf(w, "Hello this is a test");
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendData() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	t := time.Now()

	body := "hello: " + t.Format("2006-01-02 15:04:05")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}