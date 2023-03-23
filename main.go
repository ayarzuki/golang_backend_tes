package main

import (
	"goproduct/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", handlers.AddProduct).Methods("POST")
	router.HandleFunc("/api/products", handlers.ListProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
