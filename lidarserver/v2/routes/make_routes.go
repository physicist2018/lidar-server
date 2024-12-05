package routes

import (
	"net/http"

	"lidarserver.sqlc/app/lidarserver/v2/controller"
)

func NewRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /api/v1/experiments", controller.FindAll)
	r.HandleFunc("GET /api/v1/experiments/{id}", controller.FindOne)
	r.HandleFunc("DELETE /api/v1/experiments/{id}", controller.DeleteOne)
}
