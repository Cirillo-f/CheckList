package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

// [GET] /list
func GetList(w http.ResponseWriter, r *http.Request) {
	// Здесь будем хранить результат запроса
	var tasks []models.Task

	// Запрос к базе данных
	request := `SELECT * FROM tasks;`

	// Выполняем запрос
	ROWS, err := connectdb.DB.Query(request)
	if err != nil {
		log.Println("[ERROR]: Ошибка выполнения SQL-запроса:", err)
		http.Error(w, "Ошибка при получении списка задач. Попробуйте позже.", http.StatusInternalServerError)
		return
	}
	defer ROWS.Close() // Закрываем rows после завершения работы

	// Обрабатываем результаты запроса
	for ROWS.Next() {
		var t models.Task
		err := ROWS.Scan(&t.ID, &t.Title, &t.Description, &t.Status)
		if err != nil {
			log.Println("[ERROR]: Ошибка сканирования строки из базы данных:", err)
			http.Error(w, "Ошибка при обработке данных задач. Попробуйте позже.", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	// Проверяем наличие ошибок при итерации по строкам
	if err = ROWS.Err(); err != nil {
		log.Println("[ERROR]: Ошибка при чтении строк из базы данных:", err)
		http.Error(w, "Ошибка при чтении данных задач. Попробуйте позже.", http.StatusInternalServerError)
		return
	}

	// Кодируем результат в JSON и отправляем пользователю
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации ответа:", err)
		http.Error(w, "Ошибка при формировании ответа. Попробуйте позже.", http.StatusInternalServerError)
		return
	}
}
