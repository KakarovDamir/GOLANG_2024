package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", healthCheck).Methods("GET")
	r.HandleFunc("/models", models).Methods("GET")
	r.HandleFunc("/models/{id:[0-6]+}", car).Methods("GET")

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	log.Fatal(err)
}
