package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	// Декодируем запрос
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(newTask)
	if err != nil {
		log.Println("[ERROR]:Ошибка десериализации.", err)
		return
	}

	// Создаем запрос к базе данных
	request := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3);`

	//Делаем запрос
	err = connectdb.DB.Exec(request).Scan()
	if err != nil {
		log.Println("[ERROR]:Ошибка во время создания запроса.", err)
	}
}
