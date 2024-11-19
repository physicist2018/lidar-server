package routes

import (
	"fmt"
	"net/http"
)

func GetAllExperiments(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%v", []int{1, 2, 3, 4})
}
