package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {
	//Routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/Customers", getAllCustomer)
	//define a Port to server
	http.ListenAndServe("localhost:8000", nil)
}

//Write to the server and response with Welcome
func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func getAllCustomer(w http.ResponseWriter, req *http.Request) {
	customers := []Customer{
		{Name: "David", City: "San Francisco", Zipcode: "12"},
	}
	json.NewEncoder(w).Encode(customers)
}
