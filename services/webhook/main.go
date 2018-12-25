package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndPoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndPoint).Methods("POST")
	fmt.Printf("Listeing on port: 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
