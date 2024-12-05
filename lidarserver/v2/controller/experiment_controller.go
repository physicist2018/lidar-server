package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"lidarserver.sqlc/app/lidarserver/v2/services"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World!"))
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	ret := services.ES.FindAll()
	json.NewEncoder(w).Encode(ret)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	id := r.PathValue("id")

	json.NewEncoder(w).Encode(id)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	json.NewEncoder(w).Encode(vars["id"])
}
