package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	porscheapi "tsis1/pkg/porsche-api"
)

type Response struct {
	Cars []porscheapi.Car `json:"cars"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": porscheapi.Info()})
}

func models(w http.ResponseWriter, r *http.Request) {
	cars := porscheapi.GetCars()
	respondWithJSON(w, http.StatusOK, cars)
}

func car(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	restaurant, err := porscheapi.GetCar(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, restaurant)
}
