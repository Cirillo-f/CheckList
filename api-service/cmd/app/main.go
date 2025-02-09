package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Cirillo-f/CheckList/api-service/handlers"
	"github.com/Cirillo-f/CheckList/api-service/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.LogMiddleware)

	router.Get("/list", handlers.GetList)
	router.Post("/create", handlers.CreateTask)
	router.Put("/done", handlers.DoneTask)
	router.Delete("/delete", handlers.DeleteTask)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR]: Ошибка загрузки файлов .env", err)
	}
	log.Println("[SUCCES]:Апи-сервис запущен на http://localhost:" + os.Getenv("PORT"))

	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatal("[ERROR]:Ошибка запуска сервера.", err)
	}
}
