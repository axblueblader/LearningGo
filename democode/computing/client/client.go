package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func callRPC(operands, routekey string, res chan int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

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

	corrID := randomString(32)

	err = ch.Publish(
		"",       // exchange
		routekey, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          []byte(operands),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrID == d.CorrelationId {
			result, err := strconv.Atoi(string(d.Body))
			res <- result
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	operands := bodyFrom(os.Args)
	routekey := severityFrom(os.Args)
	result := make(chan int)

	log.Printf(" [x] Requesting %s on %s", routekey, operands)
	timeout := time.After(8 * time.Second)
	go callRPC(operands, routekey, result)
	var res int
waitResponse:
	for {
		select {
		case res = <-result:
			break waitResponse
		case <-timeout:
			log.Printf(" [.] Process took too long, connection timed out")
			break waitResponse
		default:
			log.Print(" [*] ")
			time.Sleep(2 * time.Second)
		}
	}

	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[2] == "" {
		s = "1 2"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "add"
	} else {
		s = os.Args[1]
	}
	return s
}
