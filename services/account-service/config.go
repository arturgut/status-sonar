package main

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-yaml/yaml.v2"
)

// Config struct
type Config struct {
	Loglevel string `yaml:"loglevel" envconfig:"ACCOUNT_SERVICE_LOGLEVEL"`

	Server struct {
		Port int `yaml:"port" envconfig:"ACCOUNT_SERVICE_SERVER_PORT"`
	} `yaml:"server"`

	Client struct {
		Host            string `yaml:"host" envconfig:"ACCOUNT_SERVICE_MONGO_HOST"`
		Port            int    `yaml:"port" envconfig:"ACCOUNT_SERVICE_MONGO_PORT"`
		Username        string `yaml:"username" envconfig:"ACCOUNT_SERVICE_MONGO_USERNAME"`
		Password        string `yaml:"password" envconfig:"ACCOUNT_SERVICE_MONGO_PASSWORD"`
		userDatabase    string `yaml:"userDatabase" envconfig:"ACCOUNT_SERVICE_MONGO_USER_DATABASE"`
		accountDatabase string `yaml:"accountDatabase" envconfig:"ACCOUNT_SERVICE_MONGO_ACCOUNT_DATABASE"`
	} `yaml:"client"`
}

var config Config // globally accessible

func loadConfiguration(filename string) {
	log.Info("Loading configuration...")
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("An error occured: ", err)
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Error("An error occured: ", err)
		panic(err)
	}
	log.Info("Configuration has been sucessfully loaded from ", filename)
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error("And error occured in readEnvVars(): Error message: \n\t", err)
	}
	log.Info("Enviroment variables have been sucessfully loaded.")
}
