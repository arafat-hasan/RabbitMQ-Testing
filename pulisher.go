package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Hello RabbitMQ")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Succesfully connected to RabbitMQ instance")

	channel, err := connection.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("TestQueue", false, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(queue)

	err = channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Succesfully published message to the queue")

}
