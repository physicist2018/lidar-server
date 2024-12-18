package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kataras/golog"
	"lidarserver.sqlc/app/lidarserver/v2/models"
	"lidarserver.sqlc/app/lidarserver/v2/services"
)

// GetAllExperiments
func GetAllExperiments(w http.ResponseWriter, r *http.Request) {
	// obtain the result
	result := services.ES.FindAll()

	// write neseccery headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// marshal result to response stream
	json.NewEncoder(w).Encode(result)
}

func GetOneExperiment(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	// obtain the result
	result := services.ES.FindOne(id)

	// write neseccery header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// marshal result to response stream
	json.NewEncoder(w).Encode(result)
}

func DeleteOneExperiment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := services.ES.DeleteOne(id)

	// write neseccery headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// marshal result to response stream
	json.NewEncoder(w).Encode(res)
}

func CreateOneExperiment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	exp := &models.ExperimentModel{}
	err := decoder.Decode(exp)
	if err != nil {
		golog.Error(err.Error())
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in JSON data"))
		return
	}

	resp := services.ES.CreateOne(exp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func UpdateOneExeriment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)
	exp := &models.ExperimentModel{}
	err := decoder.Decode(exp)
	if err != nil {
		golog.Error(err.Error())
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in JSON data"))
		return
	}

	resp := services.ES.UpdateOne(id, exp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
