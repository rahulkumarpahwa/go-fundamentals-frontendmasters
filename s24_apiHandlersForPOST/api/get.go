package api

import (
	"encoding/json"
	"fmt"
	"mod24/data"
	"net/http"
	"strconv"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id != "" {
		num_id, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		} else if num_id < len(data.GetAll()) && num_id >= 0 {
			json.NewEncoder(w).Encode(data.GetAll()[num_id])
			return
		} else {
			http.Error(w, "Not Found!", http.StatusNotFound)
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(data.GetAll())
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
