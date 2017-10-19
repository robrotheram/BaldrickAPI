package messages
import (
	"log"
	"time"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateListener(que string) {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")


	q, err := ch.QueueDeclare(
		que, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")


	go func() {
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")
		for{
			log.Printf("Checking for messages");
			for d := range msgs {
				log.Printf("Worker: %s | Received a message: %s",que, d.Body)
			}
			time.Sleep(time.Second * 5)
		}
		ch.Close()
		conn.Close()


	}()

}