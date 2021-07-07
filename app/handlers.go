package app

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

//Write to the server and response with Welcome
func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

//Function to getAllCustomers
func getAllCustomer(w http.ResponseWriter, req *http.Request) {
	customers := []Customer{
		{Name: "David", City: "San Francisco", Zipcode: "12"},
	}
	w.Header().Add("Content-type", "aplications/json")
	json.NewEncoder(w).Encode(customers)
}
