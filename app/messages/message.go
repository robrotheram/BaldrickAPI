package messages
import (
	"log"
	"time"
	"fmt"
	"strings"
	"github.com/streadway/amqp"
	"github.com/robrotheram/baldrick_engine/app/db"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func CreateChannel(channel db.Channel){
	c := db.Session.DB("YOUR_DB").C("Channels")
	if err := c.Insert(channel); err != nil {
		panic(err)
	}
	fmt.Println("Inserted a new Channel");
	CreateListener(channel.Name);
}



func findARule(rules []db.Rule, msg string) db.Rule {
	result := strings.Split(msg, " ")
	for i := range rules {
		if (result[0] == rules[i].Prifix){
			return rules[i]
		}
	}
	return db.Rule{};
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
		rules := db.GetRulesFromChannel(que); // []rules
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
				rule := findARule(rules, string(d.Body[:]))
				if ((db.Rule{}) != rule ) {
					output := rule.Process();
					log.Printf("Worker: %s | RULE PROCESS: %s",que, output)
				} else {
					log.Printf("Worker: %s | RULE PROCESS: NO RULE FOUND",que)
				}


			}
			time.Sleep(time.Second * 5)
		}
		ch.Close()
		conn.Close()
	}()

}