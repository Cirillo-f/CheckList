package models

// Структура задачи
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Завершение задачи
type DoneStatus struct {
	ID        int    `json:"id"`
	NewStatus string `json:"newstatus"`
}

// ID задачи которую мы собираемся удалить
type DeleteIDTask struct {
	ID int `json:"id"`
}
