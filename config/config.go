package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_URL      string
	DB_Username string
	DB_Password string
	DB_Name     string
	Port        int
}

var ENV *Config

func LoadConfig() {

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		fmt.Println(err.Error())
	}

}
