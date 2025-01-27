package dbrequest

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

// [DELETE] /delete
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Декодируем
	var taskID models.DeleteIDTask
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		log.Println()
		return
	}

	// Создаем переменную в которой будем хранить значение
	var task models.Task

	// Создаем запрос на получение таски которую мы хотим удалить для вывода информации
	selectRequest := `SELECT * FROM tasks where ID=$1;`
	err = connectdb.DB.QueryRow(selectRequest, taskID.ID).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если запись не найдена выполнить следующее
			log.Println("[ERROR]:Запись не найдена.", err)
			return
		}
		// Если возникла какая-либо другая ошибка
		log.Println("[ERROR]:Ошибка выполнения запроса.", err)
		return
	}

	// Составляем сообщение о том какую запись пользователь собирается удалить
	var message models.DeleteTaskMessage = models.DeleteTaskMessage{
		Text: "[SUCCES]:Задача " + task.Title + " успешно завершена",
	}

	// Создаем запрос удаления
	deleteRequest := `DELETE FROM tasks WHERE id=$1;`

	// Выполняем запрос
	_, err = connectdb.DB.Exec(deleteRequest, taskID.ID)
	if err != nil {
		log.Println("[ERROR]:Ошибка выполнения запроса.", err)
		return
	}

	// Выводим сообщение пользователю о том что задача завершена(удалена)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("[ERROR]:Ошибка десериализации.", err)
		return
	}
}
