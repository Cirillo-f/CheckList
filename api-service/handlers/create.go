package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Cirillo-f/CheckList/api-service/models"
)

// [POST] /create
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	// Декодируем тело запроса в структуру newTask
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования JSON из запроса:", err)
		http.Error(w, "Некорректный формат данных в запросе", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Сериализуем newTask в JSON
	jsonNewTask, err := json.Marshal(newTask)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации задачи в JSON:", err)
		http.Error(w, "Ошибка обработки данных задачи", http.StatusInternalServerError)
		return
	}

	// URL для создания задачи в DB-сервисе
	requestURL := "http://localhost:8081/create"

	// Отправляем POST-запрос к DB-сервису
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonNewTask))
	if err != nil {
		log.Printf("[ERROR]: Ошибка при отправке запроса к DB-сервису: %v\n", err)
		http.Error(w, "Ошибка при обращении к DB-сервису", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа от DB-сервиса
	if resp.StatusCode != http.StatusOK {
		log.Printf("[ERROR]: DB-сервис вернул ошибку: %s\n", resp.Status)
		http.Error(w, "Ошибка на стороне DB-сервиса", resp.StatusCode)
		return
	}

	// Отправляем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("[ERROR]: Ошибка при отправке ответа клиенту:", err)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
}
