package main

import (
	"log"
	"net/http"

	"github.com/Cirillo-f/CheckList/api-service/handlers"
	"github.com/Cirillo-f/CheckList/api-service/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.LogMiddleware)
	router.Get("/list", handlers.GetList)

	log.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
