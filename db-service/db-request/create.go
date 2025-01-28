package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
	_ "github.com/lib/pq"
)

// [POST] /create
func Create(w http.ResponseWriter, r *http.Request) {
	// Декодируем запрос
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования тела запроса:", err)
		http.Error(w, "Некорректный формат данных. Проверьте JSON.", http.StatusBadRequest)
		return
	}

	// Создаем запрос к базе данных
	request := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3);`

	// Выполняем запрос
	_, err = connectdb.DB.Exec(request, newTask.Title, newTask.Description, newTask.Status)
	if err != nil {
		log.Println("[ERROR]: Ошибка выполнения SQL-запроса:", err)
		http.Error(w, "Ошибка при добавлении новой задачи. Попробуйте позже.", http.StatusInternalServerError)
		return
	}

	// Создаем сообщение о новой добавленной задаче
	var message models.MessageNewTS = models.MessageNewTS{
		Text:  "[SUCCESS]: Добавлена новая задача.",
		NewTS: newTask.Title,
	}

	// Возвращаем пользователю сообщение
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации ответа:", err)
		http.Error(w, "Ошибка при формировании ответа. Попробуйте позже.", http.StatusInternalServerError)
		return
	}
}
