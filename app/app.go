//This package is used because the file is app
package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/Customers", getAllCustomer)
	//define a Port to server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
