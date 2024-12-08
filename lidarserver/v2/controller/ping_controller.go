package controller

import (
	"encoding/json"
	"net/http"

	"lidarserver.sqlc/app/lidarserver/v2/services"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	tbls := services.TS.GetTables()
	json.NewEncoder(w).Encode(tbls)
}
