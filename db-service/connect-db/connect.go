package connectdb

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR]:Ошибка во время загрузки данных из .env", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBName"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("[ERROR]:Ошибка во время подключения к базе данных", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("[ERROR]:Ошибка во время пинга к базе данных.", err)
	}

	log.Println("[SUCCES]:Успешное подключение к базе данных.")
}
