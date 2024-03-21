package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product `json:"products"`
}

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func loadPrducts() []Product {
	response, err := http.Get(productsURL + "/products")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}
	data, _ := io.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadPrducts()
	t := template.Must(template.ParseFiles("catalog.html"))
	t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	response, err := http.Get(productsURL + "/product/" + productID)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("view.html"))
	t.Execute(w, product)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ListProducts)
	mux.HandleFunc("/product/{id}", ShowProduct)
	http.ListenAndServe(":8080", mux)
}
