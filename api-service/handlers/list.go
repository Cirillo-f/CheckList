package handlers

import (
	"io"
	"log"
	"net/http"
)

// [GET] /list
func GetList(w http.ResponseWriter, r *http.Request) {
	// URL для получения списка задач из DB-сервиса
	URL := "http://localhost:8081/list"

	// Отправляем GET-запрос к DB-сервису
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("[ERROR]: Ошибка при обращении к DB-сервису:", err)
		http.Error(w, "Не удалось подключиться к DB-сервису", http.StatusInternalServerError)
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
		http.Error(w, "Не удалось отправить ответ", http.StatusInternalServerError)
		return
	}
}
