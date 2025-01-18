package main

import (
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	dbrequest "github.com/Cirillo-f/CheckList/db-service/db-request"
	"github.com/go-chi/chi/v5"
)

func main() {
	connectdb.InitDB()
	defer connectdb.DB.Close()

	dbAPP := chi.NewRouter()

	dbAPP.Get("/tasks", dbrequest.GetTasks)

	log.Println("DB-service is listening on$ http://localhost:8081")
	err := http.ListenAndServe(":8081", dbAPP)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
}
