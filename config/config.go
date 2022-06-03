package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config interface {
	GetAppName() string
}

var appConfig *Config
var configFileName string

func SetConfigFileName(fn string) {
	configFileName = fn
}

func GetYamlAppConfig() *Config {
	if appConfig == nil {
		appConfig = loadYamlConfiguration(configFileName)
	}
	return appConfig
}

func loadYamlConfiguration(fileName string) *Config {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Cannot read app configuraion : ", fileName, err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalln("Error Cannot initial app configuraion : ", fileName, err)
	}
	return &cfg
}
