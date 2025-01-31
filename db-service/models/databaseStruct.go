package models

// Структура задачи
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Структура сообщения
type MessageNewTS struct {
	Text  string `json:"text"`
	NewTS string `json:"newtask"` //New task/status
}

// Завершить задачу
type DoneStatus struct {
	ID        int    `json:"id"`
	NewStatus string `json:"newstatus"`
}

// ID задачи которую мы будем удалять
type DeleteIDTask struct {
	ID int `json:"id"`
}

// Сообщение о том, какую задачу мы удалили
type DeleteTaskMessage struct {
	Text string `json:"text"`
}
