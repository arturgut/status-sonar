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
	} `yaml:"server"`

	Client struct {
		SkipSSL bool          `yaml:"skipssl"`
		Timeout time.Duration `yaml:"timeout" envconfig:"URL_CHECKER_TIMEOUT"`
		Period  time.Duration `yaml:"period" envconfig:"URL_CHECKER_PERIOD"`
	} `yaml:"client"`

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
