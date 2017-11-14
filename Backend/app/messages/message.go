package messages
import (
	"log"
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

var q amqp.Queue;
var conn *amqp.Connection;
var ch *amqp.Channel;
var err error;

func CloseConnect()  {
	defer conn.Close()
	defer ch.Close()

}

func OpenConnect(){
	conn, err = amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err = ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

}

func CreateListener(que string) {
	rules := db.GetRulesFromChannel(que); // []rules
	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", que)
	err = ch.QueueBind(
		q.Name,        // queue name
		que,             // routing key
		"logs_direct", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
	q.Name, // queue
	"",     // consumer
	true,   // auto ack
	false,  // exclusive
	false,  // no local
	false,  // no wait
	nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			log.Printf("Worker: %s | Received a message: %s",q.Name, d.Body)
			rule := findARule(rules, string(d.Body[:]))
			if (rule.Function != "" ) {
				output := rule.Process();
				log.Printf("Worker: %s | RULE PROCESS: %s",q.Name, output)
			} else {
				log.Printf("Worker: %s | RULE PROCESS: NO RULE FOUND",q.Name)
			}
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
}
