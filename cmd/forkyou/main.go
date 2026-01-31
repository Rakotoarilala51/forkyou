package main

import (
	"fmt"
	"os"

	"github.com/Rakotoarilala51/forkyou"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}
func init() {
	rootCmd = &cobra.Command{
		Use:   "forkyou",
		Short: "Project Forking Tool for github",
	}
	rootCmd.AddCommand(forkyou.SearchCmd, forkyou.ForkCmd, forkyou.CloneCmd, forkyou.DocsCmd, forkyou.ConfigCmd)
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("lamine")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
