package routes

import (
	"net/http"

	"lidarserver.sqlc/app/lidarserver/v2/controller"
)

func NewRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", controller.Ping)
	r.HandleFunc("GET /api/v2/experiments", controller.GetAllExperiments)
	r.HandleFunc("POST /api/v2/experiments", controller.CreateOneExperiment)
	r.HandleFunc("GET /api/v2/experiments/{id}", controller.GetOneExperiment)
	r.HandleFunc("DELETE /api/v2/experiments/{id}", controller.DeleteOneExperiment)
	r.HandleFunc("PUT /api/v2/experiments/{id}", controller.UpdateOneExeriment)

	// r.HandleFunc("GET /api/v1/experiments", controller.FindAll)
	// r.HandleFunc("GET /api/v1/experiments/{id}", controller.FindOne)
	// r.HandleFunc("DELETE /api/v1/experiments/{id}", controller.DeleteOne)
}
