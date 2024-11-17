package routes

import "net/http"

func MakeRoutes(m *http.ServeMux) {
	m.Handle("GET /experiments", http.HandlerFunc(GetAllExperiments))
}
