package models

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type MessageNewTS struct {
	Text  string `json:"text"`
	NewTS string `json:"newtask"` //New task/status
}

type DoneStatus struct {
	ID        int    `json:"id"`
	NewStatus string `json:"newstatus"`
}

type DeleteIDTask struct {
	ID int `json:"id"`
}

type DeleteTaskMessage struct {
	Text string `json:"text"`
}
