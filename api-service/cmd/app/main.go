package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
