package main

import (
	"io/ioutil"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-yaml/yaml.v2"
)

// Config struct
type Config struct {
	Loglevel string `yaml:"loglevel" envconfig:"URL_CHECKER_LOGLEVEL"`

	Server struct {
		Port int `yaml:"port" envconfig:"URL_CHECKER_SERVER_PORT"`
	} `yaml:"http_server"`

	Client struct {
		SkipSSL bool          `yaml:"skipssl"`
		Timeout time.Duration `yaml:"timeout" envconfig:"URL_CHECKER_TIMEOUT"`
		Period  time.Duration `yaml:"period" envconfig:"URL_CHECKER_PERIOD"`
	} `yaml:"http_client"`

	DbClient struct {
		Host            string `yaml:"host" envconfig:"ACCOUNT_SERVICE_MONGO_HOST"`
		Port            int    `yaml:"port" envconfig:"ACCOUNT_SERVICE_MONGO_PORT"`
		Username        string `yaml:"username" envconfig:"ACCOUNT_SERVICE_MONGO_USERNAME"`
		Password        string `yaml:"password" envconfig:"ACCOUNT_SERVICE_MONGO_PASSWORD"`
		userDatabase    string `yaml:"userDatabase" envconfig:"ACCOUNT_SERVICE_MONGO_USER_DATABASE"`
		accountDatabase string `yaml:"accountDatabase" envconfig:"ACCOUNT_SERVICE_MONGO_ACCOUNT_DATABASE"`
	} `yaml:"mongo_client"`

	Urls []string
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
	log.Info("Config has been sucessfully loaded.")
	log.Debug("Config map: ", config)
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error("And error occured in readEnvVars(): Error message: \n\t", err)
	}
}
