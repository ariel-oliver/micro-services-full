package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ariel-oliver/micro-service-full/payment/queue"

	"github.com/streadway/amqp"
)

type Order struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductId string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("order_queue", connection, in)

	var order Order
	for payload := range in {
		json.Unmarshal(payload, &order)
		order.Status = "aprovado"
		notifyPaymentProcessed(order, connection)
	}
}

func notifyPaymentProcessed(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "payment_ex", "", ch)
	fmt.Println(string(json))
}
