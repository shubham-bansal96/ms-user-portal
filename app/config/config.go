package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const fileName = "config.yml"

type Configuration struct {
	Database     Database `yaml:"Database"`
	Environment  string   `yaml:"Environment"`
	MSName       string   `yaml:"MSName"`
	LogLevel     string   `yaml:"LogLevel"`
	JWTSecretKey string   `yaml:"JWTSecretKey"`
	Domain       string   `yaml:"domain"`
}

type Database struct {
	DBName   string `yaml:"DBName"`
	Password string `yaml:"Password"`
	Server   string `yaml:"Server"`
	UserName string `yaml:"Username"`
	Type     string `yaml:"Type"`
	Port     int    `yaml:"Port"`
}

var Config *Configuration

func Initialize() {
	ymlConfig, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err.Error())
	}
	if err := yaml.Unmarshal(ymlConfig, &Config); err != nil {
		log.Fatalf("Error Unmarshelling config file: %v", err.Error())
	}
}
