package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/ariel-oliver/micro-service-full/order/db"
	"github.com/ariel-oliver/micro-service-full/order/queue"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

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
	var param string
	flag.StringVar(&param, "opt", "", "Usage")
	flag.Parse()

	in := make(chan []byte)
	connection := queue.Connect()

	switch param {
	case "checkout":
		queue.StartConsuming("checkout_queue", connection, in)
		for payload := range in {
			notifyOrderCreated(createOrder(payload), connection)
			createOrder(payload)
			fmt.Println(string(payload))
		}
	case "payment":
		queue.StartConsuming("payment_queue", connection, in)
		var order Order
		for payload := range in {
			json.Unmarshal(payload, &order)
			saveOrder(order)
			fmt.Println("Payment: ", string(payload))
		}
	}

}

func createOrder(payload []byte) Order {
	var order Order
	json.Unmarshal(payload, &order)

	uuid := uuid.NewV4()
	order.Uuid = uuid.String()
	order.Status = "pendente"
	order.CreatedAt = time.Now()
	saveOrder(order)
	return order
}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)
	connection := db.Connect()
	ctx := context.Background()
	err := connection.Set(ctx, order.Uuid, string(json), 0).Err()
	if err != nil {
		panic(err.Error())
	}
	val, err := connection.Get(ctx, order.Uuid).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(order.Uuid, val)
}

func notifyOrderCreated(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "order_ex", "", ch)
}
