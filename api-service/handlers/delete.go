package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Cirillo-f/CheckList/api-service/models"
)

// [DELETE] /delete
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Создаем экземпляр переменной чтобы потом декодировать туда айди задачи
	var taskID models.DeleteIDTask

	// Декодируем и присваиваем ID задачи экземпляру структуры task_id
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}

	defer func() {
		if err = r.Body.Close(); err != nil {
			log.Println("[ERROR]: Ошибка во время закрытия соединения!", err)
		}
	}()

	// Создаем URL к которому мы будем делать DELETE запрос
	URL := os.Getenv("DB_SERVICE_URL") + "/delete"
	if URL == "" {
		URL = "http://localhost:8081" + "/delete"
	}

	// При помощи Marshal сериализуем пременную task_id
	jsonIDTask, err := json.Marshal(taskID)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации JSON:", err)
		http.Error(w, "Ошибка обработки данных", http.StatusBadRequest)
		return
	}

	// Создаем DELETE запрос
	request, err := http.NewRequest("DELETE", URL, bytes.NewBuffer(jsonIDTask))
	if err != nil {
		log.Println("[ERROR]:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок чтобы сервер к которому мы делаем запрос понимал какой тип данных ему предстоит парсить
	request.Header.Set("Content-Type", "application/json")

	// Создаем http Клиент который будет отвечать за совершение запроса к DB-сервису
	client := &http.Client{}

	// Совершаем запрос через наш новый http.Client
	resp, err := client.Do(request)
	if err != nil {
		log.Println("[ERROR]: Ошибка выполнения HTTP-запроса:", err)
		http.Error(w, "Ошибка при выполнении запроса к серверу", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println("[ERROR]: Ошибка во время закрытия соединения!", err)
		}
	}()

	// Проверяем статус ответа от сервера
	if resp.StatusCode != http.StatusOK {
		log.Printf("[ERROR]: Сервер вернул ошибку: %s\n", resp.Status)
		http.Error(w, "Ошибка на стороне сервера", resp.StatusCode)
		return
	}

	// Отправляем пользователю ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("[ERROR]: Ошибка копирования тела ответа:", err)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
}
