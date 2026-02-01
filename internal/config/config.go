package config

import (
	"os"
	"path/filepath"

	"github.com/Rakotoarilala51/forkyou/internal/filesystem"
	"github.com/spf13/viper"
)

func InitConfig() error {
	configDIr := filepath.Join(os.Getenv("HOME"), ".forkyou", "config.yaml")
	if err := filesystem.EnsureFile(configDIr); err != nil {
		return err
	}

	viper.SetConfigFile(configDIr)
	return nil
}
