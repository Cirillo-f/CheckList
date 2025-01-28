package dbrequest

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

// DB-Service: [DELETE] /delete
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Декодируем ID задачи из тела запроса
	var taskID models.DeleteIDTask
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования тела запроса:", err)
		http.Error(w, "Некорректный формат данных. Проверьте JSON.", http.StatusBadRequest)
		return
	}

	// Создаем переменную для хранения информации о задаче
	var task models.Task

	// Запрос на получение информации о задаче, которую нужно удалить
	selectRequest := `SELECT * FROM tasks WHERE ID=$1;`
	err = connectdb.DB.QueryRow(selectRequest, taskID.ID).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если запись не найдена
			log.Println("[ERROR]: Задача с указанным ID не найдена:", err)
			http.Error(w, "Задача с указанным ID не найдена.", http.StatusNotFound)
			return
		}
		// Если произошла другая ошибка
		log.Println("[ERROR]: Ошибка выполнения запроса на получение задачи:", err)
		http.Error(w, "Ошибка при получении данных задачи. Попробуйте позже.", http.StatusInternalServerError)
		return
	}

	// Формируем сообщение о задаче, которую пользователь собирается удалить
	var message models.DeleteTaskMessage = models.DeleteTaskMessage{
		Text: "[SUCCESS]: Задача \"" + task.Title + "\" успешно удалена.",
	}

	// Запрос на удаление задачи
	deleteRequest := `DELETE FROM tasks WHERE ID=$1;`
	_, err = connectdb.DB.Exec(deleteRequest, taskID.ID)
	if err != nil {
		log.Println("[ERROR]: Ошибка выполнения запроса на удаление задачи:", err)
		http.Error(w, "Ошибка при удалении задачи. Попробуйте позже.", http.StatusInternalServerError)
		return
	}

	// Возвращаем сообщение об успешном удалении
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации ответа:", err)
		http.Error(w, "Ошибка при формировании ответа. Попробуйте позже.", http.StatusInternalServerError)
		return
	}
}
