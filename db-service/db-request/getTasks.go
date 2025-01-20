package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Здесь будем хранить результат запроса
	var tasks []models.Task

	// Запрос к бд
	request := `SELECT * FROM tasks;`

	//Совершаем запрос
	ROWS, err := connectdb.DB.Query(request)
	if err != nil {
		log.Println("[ERROR]:Ошибка выполнения SQL-запроса.", err)
	}

	// Проходимся по ответу от базы данных
	for ROWS.Next() {
		var t models.Task
		err := ROWS.Scan(&t.ID, &t.Title, &t.Description, &t.Status)
		if err != nil {
			log.Println("[ERROR]:Ошибка сканирования.", err)
			return
		}
		tasks = append(tasks, t)
	}

	// Кодируем и отправляем пользователю
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Println("[ERROR]:Ошибка сериализации.")
		return
	}

}
