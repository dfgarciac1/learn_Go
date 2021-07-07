package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type is like a object
type Customer struct {
	//When put Json: this convert  the headers in full-name , city ....
	Name    string `json:"full-name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
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
	//Transform the data (header) to Json when the request is given
	w.Header().Add("Content-type", "aplications/json")
	json.NewEncoder(w).Encode(customers)
}
