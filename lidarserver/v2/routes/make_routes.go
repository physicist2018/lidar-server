package routes

import (
	"github.com/gorilla/mux"
	"lidarserver.sqlc/app/lidarserver/v2/controller"
)

func NewRoutes(r *mux.Router) {
	experimentRoutes := r.PathPrefix("/api/v1/experiments").Subrouter()
	experimentRoutes.HandleFunc("", controller.FindAll).Methods("GET")
	experimentRoutes.HandleFunc("/{id}", controller.FindOne).Methods("GET")
	experimentRoutes.HandleFunc("/{id}", controller.DeleteOne).Methods("DELETE")

}
