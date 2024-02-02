package main

import (
	"tsis1/pkg"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/cars", pkg.Cars).Methods("GET")
	router.HandleFunc("/cars/{id:[0-5]+}", pkg.CarByID).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}