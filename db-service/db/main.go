package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	dbAPP := chi.NewRouter()

	log.Println("DB-service is listening on http://localhost:8081")
	err := http.ListenAndServe(":8081", dbAPP)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
}
