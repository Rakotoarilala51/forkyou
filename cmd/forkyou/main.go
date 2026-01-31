package main

import (
	"fmt"
	"os"

	"github.com/Rakotoarilala51/forkyou/internal/cli"
	"github.com/spf13/viper"
)

func main() {
	cli.Execute()
}
func init() {
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("lamine")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
