package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("lamine")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
