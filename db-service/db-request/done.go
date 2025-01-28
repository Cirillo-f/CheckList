package dbrequest

import (
	"encoding/json"
	"log"
	"net/http"

	connectdb "github.com/Cirillo-f/CheckList/db-service/connect-db"
	"github.com/Cirillo-f/CheckList/db-service/models"
)

// [PUT] /done
func DoneTask(w http.ResponseWriter, r *http.Request) {
	// Декодируем запрос (нам нужен только новый статус и ID задачи)
	var dStatus models.DoneStatus
	err := json.NewDecoder(r.Body).Decode(&dStatus)
	if err != nil {
		log.Println("[ERROR]: Не удалось декодировать тело запроса:", err)
		http.Error(w, "Некорректный формат данных. Проверьте JSON.", http.StatusBadRequest)
		return
	}

	// Проверяем что все хорошо с параметрами
	log.Printf("SQL-Status: %s | ID: %d\n", dStatus.NewStatus, dStatus.ID)

	// Создаем запрос с изменением статуса
	request := `UPDATE tasks SET status=$1 WHERE id=$2;`

	// Выполняем запрос
	_, err = connectdb.DB.Exec(request, dStatus.NewStatus, dStatus.ID)
	if err != nil {
		log.Println("[ERROR]: Ошибка выполнения SQL-запроса:", err)
		http.Error(w, "Ошибка при обновлении статуса задачи. Попробуйте позже.", http.StatusInternalServerError)
		return
	}

	// Создаем сообщение о том, что всё прошло успешно
	var message models.MessageNewTS = models.MessageNewTS{
		Text:  "[SUCCESS]: Успешное обновление статуса.",
		NewTS: "На задачу установлен статус " + dStatus.NewStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации ответа:", err)
		http.Error(w, "Ошибка при формировании ответа. Попробуйте позже.", http.StatusInternalServerError)
		return
	}
}
