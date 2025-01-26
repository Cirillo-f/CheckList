package handlers

import (
	"io"
	"log"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	request := "http://localhost:8081/list"

	resp, err := http.Get(request)
	if err != nil {
		log.Println("[ERROR]: Failed to contact db service", err)
		http.Error(w, "Unadle to contact database service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[ERROR]: DB service return wrong status.", resp.StatusCode)
		http.Error(w, "Database service ERROR.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("[ERROR]: Failed to write response.", err)
		http.Error(w, "Failed to send response.", http.StatusInternalServerError)
		return
	}
}
