package models

type DatabaseCONFIG struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Database DatabaseCONFIG `yaml:"database"`
}

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

type UpdateStatus struct {
	ID        int    `json:"id"`
	NewStatus string `json:"newstatus"`
}

type DoneTask struct {
	ID int `json:"id"`
}

type DoneMessage struct {
	Text string `json:"text"`
}
