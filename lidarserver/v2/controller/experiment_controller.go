package controller

import (
	"encoding/json"
	"net/http"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World!"))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": "Hello World!"})
}
