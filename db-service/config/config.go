package config

import (
	"os"

	"github.com/Cirillo-f/CheckList/db-service/models"
	"gopkg.in/yaml.v3"
)

func LoadDataFromYaml(filename string) (*models.Config, error) {
	//Открываем файл
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//Парсим yaml
	var config models.Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
