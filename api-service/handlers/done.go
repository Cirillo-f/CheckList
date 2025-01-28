package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Cirillo-f/CheckList/api-service/models"
)

// [PUT] /done
func DoneTask(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи у которой мы будем менять статус
	var dStatus models.DoneStatus

	// Декодируем и присваиваем dStatus номер который получили в json
	err := json.NewDecoder(r.Body).Decode(&dStatus)
	if err != nil {
		log.Println("[ERROR]: Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Создаем URL к которому мы будем делать пост запрос
	URL := "http://localhost:8081/done"

	// При помощи Marshal сериализуем переменную dStatus
	jsonIdTask, err := json.Marshal(dStatus)
	if err != nil {
		log.Println("[ERROR]: Ошибка сериализации JSON:", err)
		http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
		return
	}

	// Создаем PUT /done запрос
	request, err := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonIdTask))
	if err != nil {
		log.Println("[ERROR]: Ошибка создания HTTP-запроса:", err)
		http.Error(w, "Ошибка при формировании запроса", http.StatusInternalServerError)
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
	defer resp.Body.Close()

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
