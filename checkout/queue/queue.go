package queue

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		dsn := "amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@" + os.Getenv("RABBITMQ_DEFAULT_HOST") + ":" + os.Getenv("RABBITMQ_DEFAULT_PORT") + os.Getenv("RABBITMQ_DEFAULT_VHOST")
		conn, err := amqp.Dial(dsn)
		if err != nil {
			fmt.Printf("Failed to connect to RabbitMQ: %s. Retrying in %s...\n", err.Error(), retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		channel, err := conn.Channel()
		if err != nil {
			fmt.Printf("Failed to open a channel: %s. Retrying in %s...\n", err.Error(), retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		return channel
	}

	panic("Failed to connect to RabbitMQ after multiple attempts")
}

func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {

	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Message sent")
}
