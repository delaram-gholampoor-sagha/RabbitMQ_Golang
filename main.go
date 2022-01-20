package main

import (
	"fmt"

	"github.com/Delaram-Gholampoor-Sagha/RabbitMQ_Golang/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("hello guys , I'm running !!!!")
	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hi I finally made it !!!!!")
	if err != nil {
		return err
	}
	app.Rmq.Consume()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error Setting Up our application")
		fmt.Println(err)
	}
}
