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
	router.Post("/create", handlers.CreateTask)
	router.Put("/done", handlers.DoneTask)
	router.Delete("/delete", handlers.DeleteTask)

	log.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("[ERROR]:Ошибка запуска сервера.", err)
	}
}
