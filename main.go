package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	log.Println("Endpoint hit: homePage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func handleRequests() {
	fmt.Println("lOGGING")
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":10000", nil)

}

func main() {
	Products = []Product{
		{Id: 1, Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: 2, Name: "Desk", Quantity: 200, Price: 200.00},
	}
	handleRequests()
}
