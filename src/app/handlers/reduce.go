package handlers

import (
	"encoding/json"
	"net/http"

	"app/models"
)

func Add(w http.ResponseWriter, r *http.Request) {
	reduce := models.Reduce{Url: "test", Token: "test"}

	js, err := json.Marshal(reduce)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Get(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()[models.KEY_TOKEN]

	if !ok || len(keys[0]) < 1 {
		http.Error(w, models.KEY_TOKEN+" param is missing", http.StatusBadRequest)
		return
	}

	reduce := models.Reduce{Url: "test", Token: keys[0]}

	js, err := json.Marshal(reduce)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
