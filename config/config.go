package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapStructure:"port"`
	Storage string `mapStructure:"storage"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	// search defined path of file
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			return nil, errors.New(".env not found")
		}
		return nil, fmt.Errorf("fatal error config file %s", err)
	}

	// unmarshal parameter file .env to struct
	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("fatal error decode : %s", err)
	}
	return &config, nil
}

type mysql struct {
	Username string
	Password string
	Host string
	Port string
	Database string
}

type configure struct {
	MySQL mysql
}

func NewConfig() configure {
	return configure{
		MySQL: mysql{
			Username: "root",
			Password: "@Ugm428660",
			Host:	  "localhost",
			Port:	  "3306",
			Database: "bootcamp",
		},
	}
}