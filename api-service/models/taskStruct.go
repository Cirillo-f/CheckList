package models

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type DoneStatus struct {
	ID        int    `json:"id"`
	NewStatus string `json:"newstatus"`
}

type DeleteIDTask struct {
	ID int `json:"id"`
}
