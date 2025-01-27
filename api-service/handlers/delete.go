package handlers

import (
	"encoding/json"
	"github.com/Cirillo-f/CheckList/api-service/models"
	"log"
	"net/http"
)

// [DELETE] /delete
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Создаем экземпляр переменной чтобы потом декодировать туда айди задачи
	var task_id models.DeleteIDTask

	// Декодируем и присваиваем ID задачи экземпляру структуры task_id
	err := json.NewDecoder(r.Body).Decode(&task_id)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//URL := "http://localhost:8081/delete"

}
