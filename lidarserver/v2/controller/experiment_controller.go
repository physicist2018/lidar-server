package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"lidarserver.sqlc/app/lidarserver/v2/services"
)

// FindAll is a function that handles HTTP requests to the "/experiments" endpoint.
// It sets the status and response header, finds all documents in the Experiment collection, encodes the result to JSON, and writes it to the response writer.
func FindAll(w http.ResponseWriter, r *http.Request) {

	// set status and response header
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	// find all documents int Experiment collection
	ret := services.ES.FindAll()

	// encode to json
	json.NewEncoder(w).Encode(ret)
}

// FindOne is a function that handles HTTP requests to the "/experiments/{id}" endpoint.
// It sets the status and response header, gets the ID from the request path, finds a document in the Experiment collection by ID, encodes the result to JSON, and writes it to the response writer.
func FindOne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")
	ret := services.ES.FindOne(id)

	json.NewEncoder(w).Encode(ret)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	json.NewEncoder(w).Encode(vars["id"])
}
