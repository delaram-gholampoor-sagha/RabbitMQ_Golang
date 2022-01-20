package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn *amqp.Connection
	channel *amqp.Channel
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to RabbitMQ")
	r.channel , err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_ ,err = r.channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	return nil
}


func (r *RabbitMQ) Publish(message string) error {
	err := r.channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)

	if err != nil {
		return err
	}
	fmt.Println("successfully published message to queue")
	return nil 
}


func (r *RabbitMQ) Consume() {
	msgs , err := r.channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,

	)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Recieved Messages: %s\n" , msg.Body)
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
