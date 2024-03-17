package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/ariel-oliver/micro-services-full/checkout/queue"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
}

var productsUrl string

func init() {
	productsUrl = os.Getenv("PRODUCT_URL")
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	response, err := http.Get(productsUrl + "/product/" + id)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductId = r.FormValue("product_id")

	data, _ := json.Marshal(order)
	fmt.Println(string(data))

	connection := queue.Connect()
	queue.Notify(data, "checkout_ex", "", connection)

	w.Write([]byte("Processou!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /finish", finish)
	mux.HandleFunc("GET /{id}", displayCheckout)
	http.ListenAndServe(":8082", mux)
}
