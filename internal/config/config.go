package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Rakotoarilala51/forkyou/internal/filesystem"
	"github.com/spf13/viper"
)

func InitConfig() error {
	configDir := filepath.Join(os.Getenv("HOME"), ".forkyou")
	configFile := filepath.Join(configDir, "config.yaml")
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	if err := filesystem.EnsureFile(configDir); err != nil {
		return err
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok || err.Error() == "While parsing config: yaml: line 1: did not find expected node content" {
			viper.Set("user.token", "")
			return viper.WriteConfig()
		}
		return err
	}
	return nil
}

func AddToken(token string) {
	viper.Set("token", token)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalln(err)
	}
}

func GetToken() string {
	token := viper.GetString("user.token")
	return token
}
