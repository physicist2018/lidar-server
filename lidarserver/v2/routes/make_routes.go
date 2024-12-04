package routes

import (
	"github.com/gorilla/mux"
	"lidarserver.sqlc/app/lidarserver/v2/controller"
)

func NewRoutes(r *mux.Router) {
	experimentRoutes := r.PathPrefix("/api/v1/experiment").Subrouter()
	experimentRoutes.HandleFunc("", controller.FindAll)
}
