package handlers

import (
	"encoding/json"
	"net/http"
	"routers/models"
)

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/help", http.StatusMovedPermanently)
}

func PostGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reqs []models.Request
	if err := json.NewDecoder(r.Body).Decode(&reqs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resps := make([]models.Response, 0, len(reqs))
	for _, req := range reqs {
		resp := models.Response{Greeting: "Привіт, " + req.Role + " " + req.Name + "!", Request: req}
		resps = append(resps, resp)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetHelp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if _, err := w.Write([]byte("Для використання сервісу відправте POST запит на /greet")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
