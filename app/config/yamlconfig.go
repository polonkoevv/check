package config

import (
	"io/ioutil"
	"library/app/models"
	"log"

	"gopkg.in/yaml.v2"
)

func GetYamlConfig(filename string) *models.Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	return &config
}
