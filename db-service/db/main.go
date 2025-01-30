package main

import (
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	dbrequest "github.com/Cirillo-f/CheckList/db-service/db-request"
	"github.com/Cirillo-f/CheckList/db-service/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	connectdb.InitDB()
	defer func() {
		if err := connectdb.DB.Close(); err != nil {
			log.Println("[ERROR]: Ошибка во закрытии соединения!", err)
		}
	}()

	dbAPP := chi.NewRouter()
	dbAPP.Use(middleware.LogMiddleware)

	dbAPP.Get("/list", dbrequest.GetList)
	dbAPP.Post("/create", dbrequest.Create)
	dbAPP.Put("/done", dbrequest.DoneTask)
	dbAPP.Delete("/delete", dbrequest.DeleteTask)

	log.Println("DB-service is listening on$ http://localhost:8081")
	err := http.ListenAndServe(":8081", dbAPP)
	if err != nil {
		log.Fatal("[ERROR]:Ошибка запуска сервера.", err)
	}
}
