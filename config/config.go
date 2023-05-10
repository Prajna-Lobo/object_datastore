package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Env   string      `mapstructure:"env"`
	Mongo MongoConfig `mapstructure:"mongo"`
}

type MongoConfig struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

func LoadConfiguration(path string) *Configuration {
	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found")
		}
		log.Fatal(err.Error())
	}

	var config Configuration
	marshalErr := viper.Unmarshal(&config)
	if marshalErr != nil {
		log.Fatal(marshalErr.Error())
	}

	return &config
}
