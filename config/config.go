package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config interface {
	GetAppName() string
}

var configFileName string

func SetConfigFileName(fn string) {
	configFileName = fn
}

var appConfig *interface{}

func GetYamlAppConfig() *interface{} {
	if appConfig == nil {
		appConfig = LoadYamlConfiguration(configFileName, appConfig)
	}
	return appConfig
}

func LoadYamlConfiguration(fileName string, appConfig *interface{}) *interface{} {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Cannot read app configuraion : ", fileName, err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(appConfig)
	if err != nil {
		log.Fatalln("Error Cannot initial app configuraion : ", fileName, err)
	}
	return appConfig
}
