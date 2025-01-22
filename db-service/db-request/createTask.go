package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
	_ "github.com/lib/pq"
)

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	// Декодируем запрос
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println("[ERROR]:Ошибка десериализации.", err)
		return
	}

	// Создаем запрос к базе данных
	request := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3);`

	//Делаем запрос
	_, err = connectdb.DB.Exec(request, newTask.Title, newTask.Description, newTask.Status)
	if err != nil {
		log.Println("[ERROR]:Ошибка во время создания запроса.", err)
		return
	}

	// Создаем сообщение о новой добавленной задаче
	var message models.MessageNewTS = models.MessageNewTS{
		Text:  "[SUCCES]:Добавлена новая задача",
		NewTS: newTask.Title,
	}

	// Возвращаем пользователю сообщение
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("[ERROR]:Ошибка сериализации.")
		return
	}
}
