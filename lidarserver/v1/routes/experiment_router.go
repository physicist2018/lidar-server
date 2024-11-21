package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"lidarserver.sqlc/app/lidarserver/v1/conrtollers"
)

func GetAllExperiments(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%v", []int{1, 2, 3, 4})
}

func GetExperimentById(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	log.Println(idstr)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid id value %v", idstr)
		return
	}

	e, err := conrtollers.GetExperimentById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request %v", err)
		return
	}

	if bts, err := json.Marshal(e); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(bts)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
