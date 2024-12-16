package routes

import (
	"net/http"

	"lidarserver.sqlc/app/lidarserver/v2/controller"
)

// NewRoutes is a function that sets up the routes for the HTTP server.
// It takes a pointer to an http.ServeMux and adds routes for handling HTTP requests.
// The routes are for handling requests to the root path, the "/api/v2/experiments" endpoint, and the "/api/v2/experiments/{id}" endpoint.
// The handlers for these routes are provided by the controller package.
func NewRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", controller.Ping)
	r.HandleFunc("GET /api/v2/experiments", controller.GetAllExperiments)
	r.HandleFunc("POST /api/v2/experiments", controller.CreateOneExperiment)
	r.HandleFunc("GET /api/v2/experiments/{id}", controller.GetOneExperiment)
	r.HandleFunc("DELETE /api/v2/experiments/{id}", controller.DeleteOneExperiment)
	r.HandleFunc("PUT /api/v2/experiments/{id}", controller.UpdateOneExeriment)
}
