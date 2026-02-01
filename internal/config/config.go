package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Rakotoarilala51/forkyou/internal/filesystem"
	"github.com/spf13/viper"
)

func InitConfig() error {
	configDIr := filepath.Join(os.Getenv("HOME"), ".forkyou", "config.yaml")
	viper.SetConfigFile(configDIr)
	viper.SetConfigType("yaml")
	viper.Set("user.token", "")
	if err := filesystem.EnsureFile(configDIr); err != nil {
		return err
	}
	viper.ReadInConfig()
	return nil
}

func AddToken(token string) {
	viper.Set("user.token", token)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalln(err)
	}
}

func GetToken() string {
	token := viper.GetString("user.token")

	return token
}
